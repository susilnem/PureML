<script lang="ts">
  import {
    Listbox,
    ListboxButton,
    ListboxOptions,
    Transition,
  } from "@rgossiaux/svelte-headlessui";
  import { ChevronDown, MoreVertical } from "lucide-svelte";
  import { tv } from "tailwind-variants";

  const selectStyles = tv({
    slots: {
      base: "focus:outline-none",
      noIcon:
        "w-full focus:outline-none flex justify-center items-center text-base text-slate-950 border border-slate-300 rounded px-4 py-2",
      onlyIcon:
        "focus:outline-none flex justify-center items-center text-base text-slate-950",
      withIcon:
        "w-full focus:outline-none flex justify-between items-center text-base text-slate-950 border border-slate-300 rounded px-4 py-2",
      list: "absolute w-32 bg-white z-50 mt-2",
    },
    variants: {
      intent: {
        primary: "text-slate-600",
        onlyIcon: "text-slate-600",
        noIcon: "text-slate-600",
      },
      fullWidth: {
        true: "w-64",
        false: "w-fit",
      },
    },
    defaultVariants: {
      intent: "primary",
      fullWidth: true,
    },
  });
  const { base, noIcon, onlyIcon, withIcon, list } = selectStyles();
  export let title: string;
  // export let onChange: (event: Event, item: string) => void = () => {};
  let selectedItem: any = [];
</script>

<div class="w-full">
  {#if $$restProps.intent === "noIcon"}
    <Listbox
      value={selectedItem}
      on:change={(e) => {
        selectedItem = e.detail;
        // onChange(e, selectedItem);
      }}
      class={base()}
    >
      <ListboxButton class={noIcon()}
        >{selectedItem.name || title}</ListboxButton
      >
      <Transition
        enter="transition duration-100 ease-out"
        enterFrom="transform scale-95 opacity-0"
        enterTo="transform scale-100 opacity-100"
        leave="transition duration-75 ease-out"
        leaveFrom="transform scale-100 opacity-100"
        leaveTo="transform scale-95 opacity-0"
      >
        <ListboxOptions class={list()}>
          <slot />
        </ListboxOptions>
      </Transition>
    </Listbox>
  {:else if $$restProps.intent === "primary"}
    <Listbox
      value={selectedItem}
      on:change={(e) => {
        selectedItem = e.detail;
        // onChange(e, selectedItem);
      }}
    >
      <ListboxButton class={withIcon()}
        >{selectedItem.name || title}
        <ChevronDown class="w-4 text-slate-950" />
      </ListboxButton>
      <Transition
        enter="transition duration-100 ease-out"
        enterFrom="transform scale-95 opacity-0"
        enterTo="transform scale-100 opacity-100"
        leave="transition duration-75 ease-out"
        leaveFrom="transform scale-100 opacity-100"
        leaveTo="transform scale-95 opacity-0"
      >
        <ListboxOptions class={list()}>
          <slot />
        </ListboxOptions>
      </Transition>
    </Listbox>
  {:else}
    <Listbox
      value={selectedItem}
      on:change={(e) => {
        selectedItem = e.detail;
        // onChange(e, selectedItem);
      }}
    >
      <ListboxButton class={onlyIcon()}>
        <MoreVertical class="w-4 text-slate-950" />
      </ListboxButton>
      <Transition
        enter="transition duration-100 ease-out"
        enterFrom="transform scale-95 opacity-0"
        enterTo="transform scale-100 opacity-100"
        leave="transition duration-75 ease-out"
        leaveFrom="transform scale-100 opacity-100"
        leaveTo="transform scale-95 opacity-0"
      >
        <ListboxOptions class={list()}>
          <slot />
        </ListboxOptions>
      </Transition>
    </Listbox>
  {/if}
</div>
