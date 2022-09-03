import { UserApi } from '~~/lib/api/user';
import { defineStore } from 'pinia';
import { useLocalStorage } from '@vueuse/core';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: useLocalStorage('pinia/auth/token', ''),
    expires: useLocalStorage('pinia/auth/expires', ''),
  }),
  getters: {
    isTokenExpired: state => {
      if (!state.expires) {
        return true;
      }

      if (typeof state.expires === 'string') {
        return new Date(state.expires) < new Date();
      }

      return state.expires < new Date();
    },
  },
  actions: {
    async logout(api: UserApi) {
      const result = await api.logout();

      if (result.error) {
        return result;
      }

      this.token = '';
      this.expires = '';

      return result;
    },
  },
});
