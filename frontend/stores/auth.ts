import { defineStore } from 'pinia';
import { useLocalStorage } from '@vueuse/core';
import { UserApi } from '~~/lib/api/user';

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
    /**
     * clearSession is used when the user cannot be logged out via the API and
     * must clear it's local session, usually when a 401 is received.
     */
    clearSession() {
      this.token = '';
      this.expires = '';
      navigateTo('/');
    },
  },
});
