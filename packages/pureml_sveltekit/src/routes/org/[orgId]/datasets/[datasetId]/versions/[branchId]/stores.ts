import { writable, type Writable } from 'svelte/store';

export interface Version {
    version: string;
    created_by: { name: string };
    logs: { key: string; data: any }[];
}

export const version1 = writable("");

export const version2 = writable("");

export const submitReviewVersion = writable("");

export const branch = writable("main");

// export const dataVersion: Writable<Record<string, Version>> = writable({});

// export const version1Logs: Writable<Record<string, string>> = writable({});

// export const version2Logs: Writable<Record<string, string>> = writable({});

// export const commonMetrics: Writable<string[]> = writable([]);