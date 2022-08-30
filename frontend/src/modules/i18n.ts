import { ViteSetupModule } from '@/types/ViteSetupModule';
import { createI18n } from 'vue-i18n';

// Import i18n resources
// https://vitejs.dev/guide/features.html#glob-import

// Don't need this? Try vitesse-lite: https://github.com/antfu/vitesse-lite
const messages = Object.fromEntries(
	Object.entries(
		import.meta.glob<{ default: any }>('../../locales/*.{y(a)?ml,json}', {
			eager: true,
		})
	).map(([key, value]) => {
		const isYamlOrJson = key.endsWith('.yaml') || key.endsWith('.json');

		return [key.slice(14, isYamlOrJson ? -5 : -4), value.default];
	})
);

export const install: ViteSetupModule = ({ app }) => {
	const i18n = createI18n({
		legacy: false,
		locale: 'en',
		messages,
		globalInjection: true,
	});

	app.use(i18n);
};
