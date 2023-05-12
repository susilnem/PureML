<script lang="ts">
  import {
    Listbox,
    ListboxButton,
    ListboxOptions,
    Transition,
  } from "@rgossiaux/svelte-headlessui";
  import { tv } from "tailwind-variants";

  const selectStyles = tv({
    slots: {
      base: "focus:outline-none",
      selectIcon:
        "focus:outline-none flex justify-center items-center text-base gap-x-2 font-medium text-slate-600 border border-slate-400 rounded px-4 py-2",
      noIcon: "focus:outline-none flex justify-center items-center",
      view: "bg-white flex justify-start rounded-lg shadow-lg",
    },
    variants: {
      intent: {
        primary: "text-slate-600",
        more: "text-slate-600",
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
  const { base, selectIcon, noIcon, view } = selectStyles();
  export let title: string;
  // export let onChange: (event: Event, item: string) => void = () => {};
  let selectedItem: any = [];
</script>

<div>
  {#if $$restProps.intent === "primary"}
    <Listbox
      value={selectedItem}
      on:change={(e) => {
        selectedItem = e.detail;
        // onChange(e, selectedItem);
      }}
      class={base()}
    >
      <ListboxButton class={selectIcon()}
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
        <ListboxOptions class="absolute w-64 bg-white z-50">
          <slot />
        </ListboxOptions>
      </Transition>
    </Listbox>
  {/if}
</div>
