import { getRealFormat, isValid } from "gtin";

// Convenient representations of gtin's format values.
enum CodeFormat {
  EAN13 = "GTIN-13", // 13-digit EAN
  UPCA = "GTIN-12", // 12-digit UPC-A
}

function isSearchShortcut(event: KeyboardEvent): boolean {
  return (event.metaKey || event.ctrlKey) && event.key === "f";
}

/**
 * For our purposes, a valid barcode fragment is any digit, with no
 * standard modifiers enabled (shift, ctrl, meta):
 */
function isCodeFragment(event: KeyboardEvent): boolean {
  return !Number.isNaN(parseInt(event.key)) && !event.ctrlKey && !event.metaKey && !event.shiftKey;
}

/**
 * Performs a naive check to ensure something is a EAN13 or UPC-A code.
 */
function isValidIshCode(code: string): boolean {
  try {
    // We check the format in advance, since we're only accepting EAN13 and UPC-A here.
    // There's no -real- reason for this limitation, besides protecting against feature
    // creep from supporting a larger number of formats.
    const realFormat = getRealFormat(code);

    return realFormat === CodeFormat.EAN13 || (realFormat === CodeFormat.UPCA && isValid(code));
  } catch (err) {
    // gtin.getRealFormat will throw on codes that are absolutely invalid,
    // as opposed to returning false for codes that look correct but fail
    // a format-specific check.
    return false;
  }
}

/**
 * Ignore code input if the user is actually trying to write anywhere:
 */
function isEventTargetInput(event: KeyboardEvent) {
  const tagName = (event.target as HTMLElement).tagName;
  return tagName != null && tagName.toUpperCase() === "INPUT";
}

/**
 * Provides utilities to:
 *
 * - Quickly open a global search utility, that forwards to the item page.
 * - Transparently listens for UPC/EAN-like barcode input, and automatically
 *   generates an appropriate search query for valid inputs; this allows consumer
 *   USB barcode scanners to be used as a quick shortcut.
 */
export function useQuickSearch() {
  const router = useRouter();

  /**
   * Tracks if the quick search dialog is active.
   *
   * We 'steal' cmd/ctrl+F from the user, but allow them to go back to
   * the browser-default search by tapping the shortcut twice.
   */
  const isActive = ref(false);

  /**
   * codeBuffer acts as intermediate buffer state for a partial
   * EAN/UPC code.
   */
  const codeBuffer = ref("");

  function onKeyDown(event: KeyboardEvent) {
    if (isSearchShortcut(event)) {
      // If quick search is already active, and the user taps cmd+F, get it out of the way
      // and allow the default browser search to kick in:
      if (isActive.value) {
        isActive.value = false;
      } else {
        isActive.value = true;
        event.preventDefault();
      }
    } else if (isCodeFragment(event) && !isEventTargetInput(event)) {
      const fragment = event.key;

      // Push this code fragment into our buffer. At this point we also
      // ensure we have no more than 13 numbers in the buffer, and if that
      // is the case, we clear it ahead of starting what we assume to be
      // a new code:
      if (codeBuffer.value.length < 13) {
        codeBuffer.value = `${codeBuffer.value}${fragment}`;
      } else {
        // Reset the buffer:
        codeBuffer.value = fragment;
      }
    } else if (event.key === "Enter" && isValidIshCode(codeBuffer.value)) {
      // If we have an active code buffer that seems valid, and the user presses Enter,
      // we want to generate a new search query from this code, as long as it seems valid.
      //
      // Consumer/most(?) USB barcode scanners will terminate valid codes with an Enter key press,
      // which is what we're reacting to here.
      const validCode = codeBuffer.value;

      // Regardless of what we do next, we also clear the code buffer here:
      codeBuffer.value = "";

      // TODO: Is there a good reason to not expose custom fields via search syntax?
      router.push({
        path: "/items",
        query: {
          q: "",
          fieldSelector: "true",
          // TODO: Barcode= is a temporary approach to support this behavior.
          fields: [encodeURIComponent(`Barcode=${validCode}`)],
        },
      });
    } else {
      // Every other key press resets the buffer - this applies to non-code values,
      // and also to pressing the Enter key without a valid code;
      codeBuffer.value = "";
    }
  }

  onMounted(() => {
    window.addEventListener("keydown", onKeyDown);
  });

  onUnmounted(() => {
    window.removeEventListener("keydown", onKeyDown);
  });

  return {
    codeBuffer,
    isActive,
  };
}
