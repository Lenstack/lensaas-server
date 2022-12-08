import {useEffect, useState} from "react";

export const ThemeToggle = () => {
    const [theme, setTheme] = useState("light")
    const [icon, setIcon] = useState("☀️")

    useEffect(() => {
        theme === "dark" ? setIcon("☀️") : setIcon("🌙")
        if (localStorage.getItem("theme") === "dark") {
            setTheme("dark")
            document.documentElement.classList.add('dark')
            return
        }
        setTheme("light")
        document.documentElement.classList.remove('dark')
    }, [theme])

    const toggleTheme = () => {
        if (theme === "light") {
            setTheme("dark")
            setIcon("🌙")
            localStorage.setItem("theme", "dark")
            document.documentElement.classList.add('dark')
        } else {
            setTheme("light")
            setIcon("☀️")
            localStorage.setItem("theme", "light")
            document.documentElement.classList.remove('dark')
        }
    }

    return (
        <div
            className={"flex items-center bg-[#2C3543] dark:bg-[#E1F2EC] rounded-full align-center p-2 cursor-pointer"}
            onClick={toggleTheme}>
            <span>{icon}</span>
        </div>
    )
}