<script lang="ts">
  import { tv } from "tailwind-variants";
  import { Box, ChevronDown, Database, File, Menu, X } from "lucide-svelte";
  import Button from "$lib/components/Button.svelte";
  import clsx from "clsx";
  import { page } from "$app/stores";
  import AvatarIcon from "$lib/components/Avatar.svelte";
  import Dropdown from "$lib/components/Dropdown.svelte";

  function linkCss(currentPage: boolean) {
    return clsx(
      currentPage ? " text-brand-200 " : " text-slate-500 ",
      " hover:text-brand-200 flex justify-center items-center px-5 cursor-pointer "
    );
  }

  const navbarStyles = tv({
    base: "navbar px-12 2xl:pr-0  max-w-screen-2xl",
    variants: {
      intent: {
        loggedIn: "",
        loggedOut: "",
      },
      fullWidth: {
        true: "w-full",
      },
    },
    defaultVariants: {
      intent: "loggedOut",
      fullWidth: true,
    },
  });

  export let open = false;
  export let user: string;
  export let orgName: string;
  export let orgAvatarName: string;

  export let orgItems = [
    { link: "/switchOrg", title: "Switch Organization" },
    { link: "/joinOrg", title: "Join Organization" },
    { link: "/createOrg", title: "Create Organization" },
  ];

  export let profileItems = [
    { link: "/profile", title: "Profile" },
    { link: "/settings", title: "Settings" },
    // { link: "/contactus", title: "Contact Us" },
    { link: "/logout", title: "Sign Out" },
  ];
</script>

{#if open}
  <ul class="bg-white rounded-b-2xl border-b border-slate-200">
    <div class="flex justify-between px-12 py-4 items-center">
      <li class="flex justify-center items-center pr-12">
        {#if $$restProps.intent === "loggedIn"}
          <AvatarIcon>{orgAvatarName}</AvatarIcon>
        {/if}
        <a href="/org/pureml" class="px-2 font-medium text-slate-600">
          {orgName}
        </a>
        {#if $$restProps.intent === "loggedIn"}
          <Dropdown intent="start">
            <p slot="trigger"><ChevronDown class="text-slate-400 w-4" /></p>
            <p slot="items">
              {#each orgItems as item}
                <li
                  class="flex px-3 py-2 text-sm justify-left items-center rounded outline-none hover:bg-slate-100 cursor-pointer"
                >
                  <a href={item.link}>
                    {item.title}
                  </a>
                </li>
              {/each}
            </p>
          </Dropdown>
        {/if}
      </li>
      <button
        on:click={() => {
          open = !open;
        }}
      >
        <X class="sm:hidden text-slate-900 cursor-pointer w-8 h-8" /></button
      >
    </div>
    <div class="px-12 py-4">
      <li class="flex items-center">
        <a href="/models"> Models </a>
      </li>
      <li class="flex items-center">
        <a href="/datasets"> Datasets </a>
      </li>
      <li class="flex items-center">
        <a href="https://pureml.mintlify.app" class="w-max"> Docs </a>
      </li>
      {#if $$restProps.intent === "loggedOut"}
        <li class="flex">
          <div
            class="w-full flex justify-center items-center px-5 font-medium text-slate-500"
          >
            <a href="/auth/signin" class="w-max"> Sign in </a>
          </div>
          <div class="w-fit">
            <Button intent="primary">Sign up</Button>
          </div>
        </li>
      {:else}
        <li class="w-full flex items-center">
          <Dropdown intent="start">
            <p slot="trigger"><AvatarIcon>{user}</AvatarIcon></p>
            <p slot="items">
              {#each profileItems as item}
                <li
                  class="flex px-3 py-2 text-sm justify-left items-center rounded outline-none hover:bg-slate-100 cursor-pointer"
                >
                  <a href={item.link}>
                    {item.title}
                  </a>
                </li>
              {/each}
            </p>
          </Dropdown>
        </li>
      {/if}
    </div>
    <li />
  </ul>
{:else}
  <div class="flex justify-center bg-slate-50 border-b border-slate-100">
    <div class={navbarStyles({ ...$$restProps })}>
      <ul class="flex w-full justify-between">
        <li class="flex justify-center items-center pr-12">
          {#if $$restProps.intent === "loggedIn"}
            <AvatarIcon intent="primary">{orgAvatarName}</AvatarIcon>
          {/if}
          <a href="/org/pureml" class="px-2 text-slate-600 font-medium">
            {orgName}
          </a>
          {#if $$restProps.intent === "loggedIn"}
            <Dropdown intent="start">
              <p slot="trigger"><ChevronDown class="text-slate-400 w-4" /></p>
              <p slot="items">
                {#each orgItems as item}
                  <li
                    class="flex px-3 py-2 text-sm justify-left items-center rounded outline-none hover:bg-slate-100 cursor-pointer"
                  >
                    <a href={item.link}>
                      {item.title}
                    </a>
                  </li>
                {/each}
              </p>
            </Dropdown>
          {/if}
        </li>
        <button
          on:click={() => {
            open = !open;
          }}
          class="md:hidden"
        >
          <Menu
            class="sm:hidden text-slate-900 cursor-pointer w-8 h-8"
          /></button
        >
        <!-- <Input
            on:click={handleMobileIconClick}
            intent="search"
          /> -->
      </ul>

      <ul class="hidden md:flex md:justify-center items-center">
        <li>
          <a
            href="/models"
            class={`${linkCss($page.url.pathname === `/models`)}`}
          >
            <Box class="w-4 h-4" />
            <span class="pl-2 font-medium">Models</span>
          </a>
        </li>

        <li>
          <a
            href="/datasets"
            class={`${linkCss($page.url.pathname === `/datasets`)}`}
          >
            <Database class="w-4 h-4" />
            <span class="pl-2 font-medium">Datasets</span>
          </a>
        </li>

        <a
          href="https://pureml.mintlify.app"
          class="flex justify-center items-center cursor-pointer px-5 hover:text-brand-200 border-r-2 border-slate-slate-200 font-medium text-slate-500"
        >
          <File class="w-4 h-4" />
          <span class="pl-2">Docs</span>
        </a>
        {#if $$restProps.intent === "loggedOut"}
          <li class="flex">
            <div
              class="w-full flex justify-center items-center px-5 font-medium text-slate-500"
            >
              <a href="/auth/signin" class="w-max"> Sign in </a>
            </div>
            <div class="w-fit">
              <Button intent="primary">Sign up</Button>
            </div>
          </li>
        {:else}
          <li class="w-full flex justify-center items-center pl-5">
            <Dropdown intent="end">
              <AvatarIcon>{user}</AvatarIcon>
              <p slot="items">
                {#each profileItems as item}
                  <li
                    class="flex px-3 py-2 text-sm justify-left items-center rounded outline-none hover:bg-slate-100 cursor-pointer"
                  >
                    <a href={item.link}>
                      {item.title}
                    </a>
                  </li>
                {/each}
              </p>
            </Dropdown>
          </li>
        {/if}
      </ul>
    </div>
  </div>
{/if}
