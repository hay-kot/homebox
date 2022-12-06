import re
import sys
from pathlib import Path

"""
This script is used in conjunction with the swagger-typescript-api NPM package.
This does some post processing on the generated typescript files to make them
more compatible with the rest of the codebase. This performs a series of regex
replacements to better align types with what the server has.

The following replacements are performed:
    1. Replaces all module precies of `Types` with empty string
    2. Replaces all optional fields with `:` instead of `?:` (due to lack of detailed swagger docs)
    3. Replaces all known date fields with `Date` instead of `string`
"""

CWD = Path(__file__).parent


def date_types(*names: list[str]) -> dict[re.Pattern, str]:
    return {re.compile(rf"{name}: string;"): rf"{name}: Date;" for name in names}


regex_replace: dict[re.Pattern, str] = {
    re.compile(r" PaginationResultRepo"): "PaginationResult",
    re.compile(r" Repo"): " ",
    re.compile(r" Services"): " ",
    re.compile(r" V1"): " ",
    re.compile(r"\?:"): ":",
    **date_types(
        "createdAt",
        "updatedAt",
        "soldTime",
        "purchaseTime",
        "warrantyExpires",
        "expiresAt",
        "date",
    ),
}


def main(args: list[str]) -> bool:
    path = Path(args[0])

    print(f"Processing {path}")

    if not path.exists():
        print(f"File {path} does not exist")
        return True

    text = "/* post-processed by ./scripts/process-types.py */\n"
    with open(path, "r") as f:
        text += f.read()

    for regex, replace in regex_replace.items():
        print(f"Replacing {regex} -> '{replace}'")
        text = regex.sub(replace, text)

    with open(path, "w") as f:
        f.write(text)

    return False


if __name__ == "__main__":
    if error := main(sys.argv[1:]):
        sys.exit(1)

    sys.exit(0)
