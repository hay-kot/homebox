<script setup lang="ts">
  import MdiGithub from "~icons/mdi/github";
  import MdiTwitter from "~icons/mdi/twitter";
  import MdiDiscord from "~icons/mdi/discord";
  import MdiFolder from "~icons/mdi/folder";

  const api = usePublicApi();

  const { data: status } = useAsyncData(async () => {
    const { data } = await api.status();
    return data;
  });
</script>
<template>
  <div>
    <AppToast />
    <div class="flex flex-col min-h-screen">
      <div class="fill-primary min-w-full absolute top-0 z-[-1]">
        <div class="bg-primary flex-col flex min-h-[20vh]" />
        <svg
          class="fill-primary drop-shadow-xl"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 1440 320"
          preserveAspectRatio="none"
        >
          <path
            fill-opacity="1"
            d="M0,32L80,69.3C160,107,320,181,480,181.3C640,181,800,107,960,117.3C1120,128,1280,224,1360,272L1440,320L1440,0L1360,0C1280,0,1120,0,960,0C800,0,640,0,480,0C320,0,160,0,80,0L0,0Z"
          ></path>
        </svg>
      </div>
      <div>
        <header class="p-4 sm:px-6 lg:p-14 sm:py-6 sm:flex sm:items-end mx-auto">
          <div>
            <h2 class="mt-1 text-4xl font-bold tracking-tight text-neutral-content sm:text-5xl lg:text-6xl flex">
              HomeB
              <AppLogo class="w-12 -mb-4" />
              x
            </h2>
            <p class="ml-1 text-lg text-base-content/50">Track, Organize, and Manage your Things.</p>
          </div>
          <div class="flex mt-6 sm:mt-0 gap-4 ml-auto text-neutral-content">
            <a class="tooltip" data-tip="Project Github" href="https://github.com/hay-kot/homebox" target="_blank">
              <MdiGithub class="h-8 w-8" />
            </a>
            <a href="https://twitter.com/haybytes" class="tooltip" data-tip="Follow The Developer" target="_blank">
              <MdiTwitter class="h-8 w-8" />
            </a>
            <a href="https://discord.gg/tuncmNrE4z" class="tooltip" data-tip="Join The Discord" target="_blank">
              <MdiDiscord class="h-8 w-8" />
            </a>
            <a href="https://hay-kot.github.io/homebox/" class="tooltip" data-tip="Read The Docs" target="_blank">
              <MdiFolder class="h-8 w-8" />
            </a>
          </div>
        </header>
        <div class="grid p-6 sm:place-items-center min-h-[50vh]">
          <slot :status="status" />
        </div>
      </div>
      <footer v-if="status" class="mt-auto text-center w-full bottom-0 pb-4">
        <p class="text-center text-sm">Version: {{ status.build.version }} ~ Build: {{ status.build.commit }}</p>
      </footer>
    </div>
  </div>
</template>
