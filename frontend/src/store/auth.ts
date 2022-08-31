import { defineStore } from 'pinia';

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
});
