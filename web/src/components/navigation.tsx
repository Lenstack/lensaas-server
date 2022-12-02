import Link from "next/link";

export const Navigation = ({children, ...restProps}: any) => {
    return (
        <nav {...restProps}>
            {children}
        </nav>
    )
}

Navigation.Item = ({children, href}: any) => {
    return (
        <Link href={href}>{children}</Link>
    )
}

Navigation.Button = ({children, ...restProps}: any) => {
    return (
        <button {...restProps}>
            {children}
        </button>
    )
}

Navigation.Container = ({children, ...restProps}: any) => {
    return (
        <div {...restProps}>
            {children}
        </div>
    )
}