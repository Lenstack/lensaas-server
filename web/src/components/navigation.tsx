import Link from "next/link";

export const Navigation = ({children, ...restProps}: any) => {
    return (
        <nav className="flex justify-between items-center py-3" {...restProps}>
            {children}
        </nav>
    )
}

Navigation.Item = ({children, href}: any) => {
    return (
        <Link href={href} className="text-lg font-bold text-red-600">{children}</Link>
    )
}