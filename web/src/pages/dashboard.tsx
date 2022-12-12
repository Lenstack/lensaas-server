import {DashboardAside, DashboardContent, DashboardHeader} from "@/containers";

export default function Dashboard() {
    return (
        <div className="min-h-screen overflow-hidden grid grid-cols-12 auto-rows-auto gap-5">
            <DashboardHeader/>
            <DashboardAside/>
            <DashboardContent/>
        </div>
    )
}