export const PageType = {
    Private: 'private',
    Public: 'public',
} as const;
export type PageType = typeof PageType[keyof typeof PageType];
export const AllPageType = Object.values(PageType);