<script lang="ts">
  import { tv } from "tailwind-variants";
  import { Box, ChevronDown, Database, File, Menu, X } from "lucide-svelte";
  import Button from "$lib/components/Button.svelte";
  import clsx from "clsx";
  import { page } from "$app/stores";
  import Avatar from "$lib/components/Avatar.svelte";
  import Dropdown from "$lib/components/Dropdown.svelte";
  import logo from "$lib/logos/PureMLLogoText.svg";
  import { base } from '$app/paths';

  function linkCss(currentPage: boolean) {
    return clsx(
      currentPage ? " text-brand-200 font-medium " : " text-slate-500 ",
      " hover:text-brand-200 flex justify-center items-center px-5 cursor-pointer "
    );
  }

  const navbarStyles = tv({
    base: "navbar px-12 py-0 max-w-screen-2xl",
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
    { link: "/auth/signout", title: "Sign Out" },
  ];
</script>

{#if open}
  <ul class="bg-white rounded-b-2xl border-b border-slate-200">
    <div class="flex justify-between px-12 py-4 items-center">
      {#if $$restProps.intent === "loggedOut"}
        <li><img src={logo} alt="Logo" class="h-8" /></li>
      {:else}
        <li class="flex justify-center items-center pr-12">
          <Avatar>{orgAvatarName || "U"}</Avatar>
          <div class="px-2 font-medium text-slate-600">
            <slot name="orgName" />
          </div>
          <Dropdown intent="start">
            <p slot="trigger"><ChevronDown class="text-slate-400 w-4" /></p>
            <p slot="items">
              {#each orgItems as item}
                <li
                  class="flex px-3 py-2 text-sm justify-left items-center rounded outline-none hover:bg-slate-100 cursor-pointer"
                >
                  <a href={item.link} class="w-full">
                    {item.title}
                  </a>
                </li>
              {/each}
            </p>
          </Dropdown>
        </li>
      {/if}
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
        <a href="{base}/models"> Models </a>
      </li>
      <li class="flex items-center">
        <a href="{base}/datasets"> Datasets </a>
      </li>
      <li class="flex items-center">
        <a href="https://pureml.mintlify.app" class="w-max"> Docs </a>
      </li>
      {#if $$restProps.intent === "loggedOut"}
        <li class="flex">
          <div
            class="w-full flex justify-center items-center px-5 font-medium text-slate-500"
          >
            <a href="{base}/auth/signin" class="w-max"> Sign in </a>
          </div>
          <div class="w-fit">
            <Button intent="primary">Sign up</Button>
          </div>
        </li>
      {:else}
        <li class="w-full flex items-center">
          <Dropdown intent="start">
            <p slot="trigger"><Avatar>{user || "U"}</Avatar></p>
            <p slot="items">
              {#each profileItems as item}
                <li
                  class="flex px-3 py-2 text-sm justify-left items-center rounded outline-none hover:bg-slate-100 cursor-pointer"
                >
                  <a href={item.link} class="w-full">
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
  <div class="flex justify-center bg-slate-50 border-b border-slate-200">
    <div class={navbarStyles({ ...$$restProps })}>
      <ul class="flex w-full justify-between h-full">
        {#if $$restProps.intent === "loggedOut"}
          <li><img src={logo} alt="Logo" class="h-8" /></li>
        {:else}
          <li class="flex justify-center items-center pr-12 h-full">
            <div class="flex items-center h-full">
              <Avatar intent="primary">{orgAvatarName || "U"}</Avatar>
              <div class="px-2 text-slate-600">
                <slot name="orgName" />
              </div>
              <Dropdown intent="start">
                <p slot="trigger"><ChevronDown class="text-slate-400 w-4" /></p>
                <p slot="items">
                  {#each orgItems as item}
                    <li
                      class="flex px-3 py-2 text-sm justify-left items-center rounded outline-none hover:bg-slate-100 cursor-pointer"
                    >
                      <a href={item.link} class="w-full">
                        {item.title}
                      </a>
                    </li>
                  {/each}
                </p>
              </Dropdown>
            </div>
          </li>
        {/if}

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
      </ul>

      <ul class="hidden md:flex md:justify-center items-center">
        <li>
          <a
            href="{base}/models"
            class={`${linkCss($page.url.pathname === `/models`)}`}
          >
            <Box class="w-4 h-4" />
            <span class="pl-2">Models</span>
          </a>
        </li>
        <li>
          <a
            href="{base}/datasets"
            class={`${linkCss($page.url.pathname === `/datasets`)}`}
          >
            <Database class="w-4 h-4" />
            <span class="pl-2">Datasets</span>
          </a>
        </li>
        <li>
          <a
            href="https://pureml.mintlify.app"
            class="flex justify-center items-center cursor-pointer px-5 hover:text-brand-200 text-slate-500"
          >
            <File class="w-4 h-4" />
            <span class="pl-2">Docs</span>
          </a>
        </li>
        {#if $$restProps.intent === "loggedOut"}
          <li class="flex">
            <div
              class="w-full flex justify-center items-center px-5 font-medium text-slate-500"
            >
              <a href="{base}/auth/signin" class="w-max"> Sign in </a>
            </div>
            <div class="w-fit">
              <Button intent="primary">Sign up</Button>
            </div>
          </li>
        {:else}
          <!-- <Input
            on:click={handleMobileIconClick}
            intent="search"
          /> -->
          <li class="w-full flex justify-center items-center pl-5">
            <Dropdown intent="end">
              <Avatar intent="large">{user || "U"}</Avatar>
              <p slot="items">
                {#each profileItems as item}
                  <li
                    class="flex px-3 py-2 text-sm justify-left items-center rounded outline-none hover:bg-slate-100 cursor-pointer"
                  >
                    <a href={item.link} class="w-full">
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
