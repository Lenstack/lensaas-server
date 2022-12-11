import {useEffect, useState} from "react";
import {useTheme} from "next-themes";

export const ThemeToggle = () => {
    const [icon, setIcon] = useState("")
    const {theme, systemTheme, resolvedTheme, setTheme} = useTheme()

    useEffect(() => {
        switch (theme) {
            case "dark":
                setIcon("☀️")
                break
            case "light":
                setIcon("🌙")
                break
            default:
                setIcon(systemTheme === "dark" ? "☀️" : "🌙")
                break
        }
    }, [systemTheme, theme])

    const toggleTheme = () => {
        resolvedTheme === "dark" ? setTheme("light") : setTheme("dark")
        resolvedTheme === "dark" ? setIcon("☀️") : setIcon("🌙")
    }

    return (
        <div
            className={"flex items-center bg-[#2C3543] dark:bg-[#E1F2EC] rounded-full align-center p-2 cursor-pointer"}
            onClick={toggleTheme}>
            <span>{icon}</span>
        </div>
    )
}