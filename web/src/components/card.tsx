export const Card = ({children, ...restProps}: any) => {
    return (
        <div {...restProps}>
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