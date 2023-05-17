## Setup

```bash
cd pureml_sveltekit
pnpm install
```

## Running server

```bash
pnpm run dev
```

## Structure

```
.
├── README.md
├── package-lock.json
├── package.json
├── postcss.config.cjs
├── src
│   ├── app.d.ts
│   ├── app.html
│   ├── app.postcss
│   ├── lib
│   │   ├── components
│   │   │   ├── Avatar.svelte
│   │   │   ├── Breadcrumbs.svelte
│   │   │   ├── Button.svelte
│   │   │   ├── Card.svelte
│   │   │   ├── Comparision.svelte
│   │   │   ├── Dropdown.svelte
│   │   │   ├── Input.svelte
│   │   │   ├── Link.svelte
│   │   │   ├── Modal.svelte
│   │   │   ├── Navbar.svelte
│   │   │   ├── Select.svelte
│   │   │   ├── Tabbar.svelte
│   │   │   └── Tag.svelte
│   │   ├── images
│   │   │   └── error
│   │   │       ├── CLILoginError.svg
│   │   │       ├── ErrorFunction.gif
│   │   │       └── ErrorNotFound404.gif
│   │   └── logos
│   │       ├── PureMLLogo.svg
│   │       ├── PureMLLogoText.svg
│   │       └── PureMLMetaTagImg.svg
│   └── routes
│       ├── +error.svelte
│       ├── +layout.server.ts
│       ├── +layout.svelte
│       ├── +page.server.ts
│       ├── +page.svelte
│       ├── +page.ts
│       ├── [username]
│       │   ├── +page.server.ts
│       │   └── +page.svelte
│       ├── api
│       │   ├── auth.server.ts
│       │   ├── datasets.server.ts
│       │   ├── models.server.ts
│       │   └── org.server.ts
│       ├── auth
│       │   ├── +layout.svelte
│       │   ├── +page.svelte
│       │   ├── +page.ts
│       │   ├── forgot_password
│       │   │   └── +page.svelte
│       │   ├── signin
│       │   │   ├── +page.server.ts
│       │   │   └── +page.svelte
│       │   ├── signout
│       │   │   ├── +page.server.ts
│       │   │   └── +page.svelte
│       │   └── signup
│       │       └── +page.svelte
│       ├── datasets
│       │   ├── +page.server.ts
│       │   ├── +page.svelte
│       │   └── +page.ts
│       ├── models
│       │   ├── +page.server.ts
│       │   ├── +page.svelte
│       │   └── +page.ts
│       ├── org
│       │   ├── +page.svelte
│       │   └── [orgId]
│       │       ├── +error.svelte
│       │       ├── +page.server.ts
│       │       ├── +page.svelte
│       │       ├── datasets
│       │       │   ├── +page.server.ts
│       │       │   ├── +page.svelte
│       │       │   └── [datasetId]
│       │       │       ├── +page.server.ts
│       │       │       ├── +page.svelte
│       │       │       ├── review
│       │       │       │   ├── +layout.server.ts
│       │       │       │   ├── +layout.svelte
│       │       │       │   ├── +page.svelte
│       │       │       │   ├── [reviewId]
│       │       │       │   │   ├── datalineage
│       │       │       │   │   │   └── +page.svelte
│       │       │       │   │   ├── graphs
│       │       │       │   │   │   └── +page.svelte
│       │       │       │   │   └── rejected
│       │       │       │   │       └── +page.svelte
│       │       │       │   ├── approved
│       │       │       │   │   └── +page.svelte
│       │       │       │   └── rejected
│       │       │       │       └── +page.svelte
│       │       │       └── versions
│       │       │           ├── +page.svelte
│       │       │           └── [branchId]
│       │       │               ├── +layout.server.ts
│       │       │               ├── +layout.svelte
│       │       │               ├── datalineage
│       │       │               │   └── +page.svelte
│       │       │               ├── graphs
│       │       │               │   └── +page.svelte
│       │       │               └── stores.ts
│       │       └── models
│       │           ├── +page.server.ts
│       │           ├── +page.svelte
│       │           └── [modelId]
│       │               ├── +page.server.ts
│       │               ├── +page.svelte
│       │               ├── review
│       │               │   ├── +layout.server.ts
│       │               │   ├── +layout.svelte
│       │               │   ├── +page.svelte
│       │               │   ├── [reviewId]
│       │               │   │   ├── graphs
│       │               │   │   │   └── +page.svelte
│       │               │   │   ├── logs
│       │               │   │   │   └── +page.svelte
│       │               │   │   └── rejected
│       │               │   │       └── +page.svelte
│       │               │   ├── approved
│       │               │   │   └── +page.svelte
│       │               │   └── rejected
│       │               │       └── +page.svelte
│       │               └── versions
│       │                   ├── +page.svelte
│       │                   └── [branchId]
│       │                       ├── +layout.server.ts
│       │                       ├── +layout.svelte
│       │                       ├── graphs
│       │                       │   └── +page.svelte
│       │                       ├── performance
│       │                       │   └── +page.svelte
│       │                       └── stores.ts
│       ├── profile
│       │   └── +page.svelte
│       ├── reset-password
│       │   └── +page.svelte
│       └── settings
│           ├── +error.svelte
│           ├── +layout.server.ts
│           ├── +layout.svelte
│           ├── +page.server.ts
│           ├── +page.svelte
│           ├── apitokens
│           │   ├── +page.server.ts
│           │   └── +page.svelte
│           ├── members
│           │   ├── +page.server.ts
│           │   └── +page.svelte
│           └── organization
│               ├── +page.server.ts
│               └── +page.svelte
├── static
│   ├── favicon.svg
│   └── robots.txt
├── svelte.config.js
├── tailwind.config.cjs
├── tsconfig.json
└── vite.config.ts
```

## Installed Packages

- [x] TailwindCSS
- [x] DaisyUI
- [x] Tailwind forms
- [x] typography
- [x] Svelvet
