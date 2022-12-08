import {useState} from "react";

export const ThemeToggle = () => {
    const [theme, setTheme] = useState("light");
    const [icon, setIcon] = useState("🌙");

    const toggleTheme = () => {
        if (theme === "light") {
            setTheme("dark");
            setIcon("☀️");
            document.documentElement.setAttribute("class", "dark");
            localStorage.setItem("theme", "light");
        } else {
            setTheme("light");
            setIcon("🌙");
            document.documentElement.setAttribute("class", "light");
            localStorage.setItem("theme", "dark");
        }
    }

    return (
        <div className={"flex items-center bg-[radial-gradient(ellipse_at_bottom,_var(--tw-gradient-stops))] from-teal-700 via-emerald-700 to-slate-400 rounded-full align-center p-2 cursor-pointer"}
             onClick={toggleTheme}>
            <span>{icon}</span>
        </div>
    )
}