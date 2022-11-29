import Link from "next/link";

export default function Navigation({children}: any) {
    return (
        <nav>
            <ul>
                {children}
            </ul>
        </nav>
    )
}

Navigation.Item = function NavigationItem({href, children}: any) {
    return (
        <li>
            <Link href={href}> {children}</Link>
        </li>
    )
}