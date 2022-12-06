import Link from "next/link";

export const Navigation = ({children, ...restProps}: any) => {
    return (
        <nav {...restProps}>
            {children}
        </nav>
    )
}

Navigation.Item = ({children, href, ...restProps}: any) => {
    return (
        <Link href={href} {...restProps}>{children}</Link>
    )
}

Navigation.Container = ({children, ...restProps}: any) => {
    return (
        <div {...restProps}>
            {children}
        </div>
    )
}