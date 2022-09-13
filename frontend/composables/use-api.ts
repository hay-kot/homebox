import { PublicApi } from "~~/lib/api/public";
import { UserApi } from "~~/lib/api/user";
import { Requests } from "~~/lib/requests";
import { useAuthStore } from "~~/stores/auth";

export type Observer = {
  handler: (r: Response) => void;
};

export type RemoveObserver = () => void;

const observers: Record<string, Observer> = {};

export function defineObserver(key: string, observer: Observer): RemoveObserver {
  observers[key] = observer;

  return () => {
    delete observers[key];
  };
}

function logger(r: Response) {
  console.log(`${r.status}   ${r.url}   ${r.statusText}`);
}

export function usePublicApi(): PublicApi {
  const requests = new Requests("", "", {});
  return new PublicApi(requests);
}

export function useUserApi(): UserApi {
  const authStore = useAuthStore();

  const requests = new Requests("", () => authStore.token, {});
  requests.addResponseInterceptor(logger);
  requests.addResponseInterceptor(r => {
    if (r.status === 401) {
      authStore.clearSession();
    }
  });

  for (const [_, observer] of Object.entries(observers)) {
    requests.addResponseInterceptor(observer.handler);
  }

  return new UserApi(requests);
}
