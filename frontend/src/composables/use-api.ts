import { PublicApi } from '@/api/public';
import { UserApi } from '@/api/user';
import { Requests } from '@/lib/requests';
import { useAuthStore } from '@/store/auth';

async function ApiDebugger(r: Response) {
	console.table({
		'Request Url': r.url,
		'Response Status': r.status,
		'Response Status Text': r.statusText,
	});
}

export function usePublicApi(): PublicApi {
	const requests = new Requests('', '', {}, ApiDebugger);
	return new PublicApi(requests);
}

export function useUserApi(): UserApi {
	const authStore = useAuthStore();
	const requests = new Requests('', () => authStore.token, {}, ApiDebugger);
	return new UserApi(requests);
}
