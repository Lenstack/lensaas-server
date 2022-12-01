import Link from "next/link";

export const Navigation = ({children, ...restProps}: any) => {
    return (
        <nav className="" {...restProps}>
            {children}
        </nav>
    )
}

Navigation.Item = ({children, href}: any) => {
    return (
        <Link href={href} className="">{children}</Link>
    )
}

Navigation.Container = ({children, ...restProps}: any) => {
    return (
        <div className="" {...restProps}>
            {children}
        </div>
    )
}