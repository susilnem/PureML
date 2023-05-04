<script lang="ts">
  import type { VariantProps } from "tailwind-variants";
  import { tv } from "tailwind-variants";

  const avatarStyles = tv({
    base: "avatar flex items-center justify-center text-blue-600 rounded-full",
    variants: {
      intent: {
        primary: "bg-blue-200 w-6",
        profile: "w-6",
        org: "bg-blue-200 w-9 h-9",
        avatar: "bg-blue-200 w-8 h-8 text-base",
      },
    },
    defaultVariants: {
      intent: "primary",
      fullWidth: true,
    },
  });

  interface $$Props extends VariantProps<typeof avatarStyles> {}
</script>

<div class="h-full">
  <div class={avatarStyles({ ...$$restProps })}>
    <slot />
  </div>
</div>
