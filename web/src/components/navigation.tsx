import Link from "next/link";

export const Navigation = ({children, ...restProps}: any) => {
    return (
        <nav className="container mx-auto" {...restProps}>
            {children}
        </nav>
    )
}

Navigation.Item = ({children, href}: any) => {
    return (
        <Link href={href} className="no-underline">{children}</Link>
    )
}

Navigation.Button = ({children, ...restProps}: any) => {
    return (
        <button className="" {...restProps}>
            {children}
        </button>
    )
}

Navigation.Container = ({children, ...restProps}: any) => {
    return (
        <div className="flex gap-5" {...restProps}>
            {children}
        </div>
    )
}