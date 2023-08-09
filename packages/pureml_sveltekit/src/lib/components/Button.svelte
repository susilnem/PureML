<script lang="ts">
  import type { VariantProps } from "tailwind-variants";
  import { tv } from "tailwind-variants";

  const button = tv({
    base: "font-medium",
    variants: {
      intent: {
        primary: "btn btn-primary btn-sm text-white",
        secondary:
          "btn btn-secondary btn-sm bg-transparent text-slate-600 border border-slate-400",
        danger: "btn btn-error btn-sm",
        icon: "btn btn-secondary btn-sm bg-transparent text-slate-600 border border-slate-400",
        org: "btn btn-secondary text-slate-600",
      },
        fullWidth: {
        true: "w-full",
        false: "w-inherit",
      },
    },
    defaultVariants: {
      intent: "primary",
      fullWidth: true,
    },
  });

  interface $$Props extends VariantProps<typeof button> {}
</script>

<button class={button({ ...$$restProps })}>
  <div class="flex items-center w-max">
    <slot />
  </div>
</button>
