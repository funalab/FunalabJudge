export const RoleType = {
    Admin: 'admin',
    Manager: 'manager',
    User: 'user',
} as const;
export type RoleType = typeof RoleType[keyof typeof RoleType];
export const AllRoleType = Object.values(RoleType);