import {Pricing} from "@/components";
import {PLAN_CONTENT} from "@/constants";

export const HomePricing = () => {
    return (
        <section className="container mx-auto p-24 flex flex-col gap-10">
            <section className={"flex flex-col gap-10 text-center"}>
                <h2 className={"text-5xl font-bold text-center"}>Pricing plans</h2>
                <p className={"text-2xl"}>Lorem ipsum dolor sit amet, consectetur adipisicing elit.
                    Aperiam doloremque laudantium modi pariatur reprehenderit!</p>
            </section>
            <Pricing className="grid grid-cols-4 gap-10">
                {
                    PLAN_CONTENT.map(({id, title, description, price, features}) => (
                        <Pricing.Item key={id} title={title} description={description} price={price}
                                      features={features}/>
                    ))
                }
            </Pricing>
        </section>
    )
}