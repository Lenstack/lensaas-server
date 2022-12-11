export const Button = ({children, ...restProps}: any) => {
    return (
        <button {...restProps} className="py-5 px-10 rounded-full bg-zinc-800 text-white dark:bg-zinc-800">
            {children}
        </button>
    )
}