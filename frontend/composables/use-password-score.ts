import type { ComputedRef, Ref } from "vue";
import { scorePassword } from "~~/lib/passwords";

export interface PasswordScore {
  score: ComputedRef<number>;
  message: ComputedRef<string>;
  isValid: ComputedRef<boolean>;
}

export function usePasswordScore(pw: Ref<string>, min = 30): PasswordScore {
  const score = computed(() => {
    return scorePassword(pw.value) || 0;
  });

  const message = computed(() => {
    if (score.value < 20) {
      return "Very weak";
    } else if (score.value < 40) {
      return "Weak";
    } else if (score.value < 60) {
      return "Good";
    } else if (score.value < 80) {
      return "Strong";
    }
    return "Very strong";
  });

  const isValid = computed(() => {
    return score.value >= min;
  });

  return {
    score,
    isValid,
    message,
  };
}
