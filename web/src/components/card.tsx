export const Card = ({children, ...restProps}: any) => {
    return (
        <div {...restProps}
             className={"bg-[#F6F6F6] dark:bg-zinc-800 dark:text-white rounded-2xl p-10 flex flex-col gap-5 text-center"}>
            {children}
        </div>
    )
}

Card.Header = ({children, ...restProps}: any) => {
    return (
        <div {...restProps}>
            {children}
        </div>
    )
}

Card.Content = ({children, ...restProps}: any) => {
    return (
        <div {...restProps}>
            {children}
        </div>
    )
}

Card.Footer = ({children, ...restProps}: any) => {
    return (
        <div {...restProps}>
            {children}
        </div>
    )
}