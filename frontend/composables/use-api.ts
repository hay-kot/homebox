import { PublicApi } from '~~/lib/api/public';
import { UserApi } from '~~/lib/api/user';
import { Requests } from '~~/lib/requests';
import { useAuthStore } from '~~/stores/auth';

function logger(r: Response) {
  console.log(`${r.status}   ${r.url}   ${r.statusText}`);
}

export function usePublicApi(): PublicApi {
  const requests = new Requests('', '', {});
  return new PublicApi(requests);
}

export function useUserApi(): UserApi {
  const authStore = useAuthStore();

  const requests = new Requests('', () => authStore.token, {});
  requests.addResponseInterceptor(logger);
  requests.addResponseInterceptor(r => {
    if (r.status === 401) {
      authStore.clearSession();
    }
  });

  return new UserApi(requests);
}
