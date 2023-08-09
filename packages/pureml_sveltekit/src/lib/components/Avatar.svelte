<script lang="ts">
  import type { VariantProps } from "tailwind-variants";
  import { tv } from "tailwind-variants";

  const avatarStyles = tv({
    base: "avatar flex items-center justify-center text-blue-600 rounded-full",
    variants: {
      intent: {
        primary: "bg-blue-200 w-6",
        large: "bg-blue-200 w-9 h-9",
        unfilled: "w-full h-full",
      },
    },
    defaultVariants: {
      intent: "primary",
      fullWidth: true,
    },
  });

  interface $$Props extends VariantProps<typeof avatarStyles> {}
</script>

<div class="flex h-full items-center">
  <div class={avatarStyles({ ...$$restProps })}>
    <slot />
  </div>
</div>
