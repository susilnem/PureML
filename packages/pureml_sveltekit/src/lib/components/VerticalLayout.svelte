<div
  class="flex justify-center sticky top-0 bg-slate-50 w-full border-b border-slate-200"
>
  <div class="flex px-12 2xl:pr-0 w-full max-w-screen-2xl">
    <slot name="tabbar" />
  </div>
</div>
<div class="flex justify-center w-full">
  <div
    class="bg-slate-50 flex flex-col h-screen overflow-hidden w-full 2xl:max-w-screen-2xl"
  >
    <div class="flex justify-between h-full">
      <div
        class="px-12 pt-8 w-1/4 max-w-1/5 bg-slate-50 border-r-2 border-slate-200"
      >
        <slot name="sidebar" />
      </div>
      <slot name="content" />
    </div>
  </div>
</div>
