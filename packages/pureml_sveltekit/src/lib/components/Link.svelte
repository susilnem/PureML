<script lang="ts">
  import { tv, type VariantProps } from "tailwind-variants";

  const linkStyles = tv({
    base: "text-brand-200 hover:text-blue-250 font-medium text-brand-200",
    variants: {
      intent: {
        primary: "text-base w-fit h-fit",
        secondary: "",
      },
    },
    defaultVariants: {
      intent: "primary",
    },
  });

  export let hyperlink: string;
</script>

<a href={hyperlink} class={linkStyles({ ...$$restProps })}>
  <slot />
</a>
