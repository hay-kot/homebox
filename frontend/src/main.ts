import App from '@/App.vue';
import { ViteSSG } from 'vite-ssg';

import '@/styles/index.css';
import { ViteSetupModule } from './types/ViteSetupModule';
import { extendedRoutes } from '@/router';

export const createApp = ViteSSG(
	App,
	{ routes: extendedRoutes },
	async ctx => {
		Object.values(
			import.meta.glob<{ install: ViteSetupModule }>('./modules/*.ts', {
				eager: true,
			})
		).map(i => i.install?.(ctx));
	},
	{}
);
