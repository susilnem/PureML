<script lang="ts">
  import { Search } from "lucide-svelte";
  import { tv } from "tailwind-variants";

  const inputStyles = tv({
    base: "focus:outline-0 focus:outline-offset-0 focus:ring-0",
    variants: {
      intent: {
        primary: "input input-primary input-sm text-slate-600",
        search:
          "input input-secondary !border-slate-200 !border flex items-center",
      },
      fullWidth: {
        true: "w-full",
        false: "w-[24rem]",
      },
      type: {
        text: "text",
        password: "password",
        email: "email",
        number: "number",
      },
    },
    defaultVariants: {
      intent: "primary",
      fullWidth: true,
      type: "text",
    },
  });

  export let type: string;
</script>

{#if $$restProps.intent === "primary"}
  <input {type} {...$$restProps} class={inputStyles({ ...$$restProps })} />
{:else}
  <div class={inputStyles({ ...$$restProps })}>
    <Search class="w-4 h-4" />
    <input
      {type}
      {...$$restProps}
      class="border-none focus:outline-none pl-2 w-full"
    />
  </div>
{/if}
