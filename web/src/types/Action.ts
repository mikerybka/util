export default interface Action {
    primary?: boolean;
    text: string;
    href?: string;
    onClick?: () => Promise<void>;
}
