// https://gist.github.com/ucw/67f7291c64777fb24341e8eae72bcd24
import type { IncomingMessage } from "http";
import type internal from "stream";
import { defineNuxtModule, logger } from "@nuxt/kit";
// Related To
// - https://github.com/nuxt/nuxt/issues/15417
// - https://github.com/nuxt/cli/issues/107
//
// fix from
// - https://gist.github.com/ucw/67f7291c64777fb24341e8eae72bcd24
// eslint-disable-next-line
import { createProxyServer } from "http-proxy";

export default defineNuxtModule({
  defaults: {
    target: "ws://localhost:7745",
    path: "/api/v1/ws",
  },
  meta: {
    configKey: "websocketProxy",
    name: "Websocket proxy",
  },
  setup(resolvedOptions, nuxt) {
    if (!nuxt.options.dev || !resolvedOptions.target) {
      return;
    }

    nuxt.hook("listen", server => {
      const proxy = createProxyServer({
        ws: true,
        secure: false,
        changeOrigin: true,
        target: resolvedOptions.target,
      });

      const proxyFn = (req: IncomingMessage, socket: internal.Duplex, head: Buffer) => {
        if (req.url && req.url.startsWith(resolvedOptions.path)) {
          proxy.ws(req, socket, head);
        }
      };

      server.on("upgrade", proxyFn);

      nuxt.hook("close", () => {
        server.off("upgrade", proxyFn);
        proxy.close();
      });

      logger.info(`Websocket dev proxy started on ${resolvedOptions.path}`);
    });
  },
});
