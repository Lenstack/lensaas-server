import {Card} from "@/components";

export const HomePricing = () => {
    return (
        <section className="container mx-auto p-24 flex flex-col gap-10">
            <section className={"flex flex-col gap-10 text-center"}>
                <h2 className={"text-5xl font-bold text-center"}>Pricing plans</h2>
                <p className={"text-2xl text-gray-700"}>Lorem ipsum dolor sit amet, consectetur adipisicing elit.
                    Aperiam doloremque laudantium modi pariatur reprehenderit!</p>
            </section>
            <section className="grid grid-cols-3 gap-10">
                <Card className="bg-white rounded-2xl p-10">
                    <Card.Header>
                        <h2 className="text-2xl font-bold">Free</h2>
                    </Card.Header>
                    <Card.Content>
                        <p className="text-gray-500">Lorem ipsum dolor sit amet, consectetur adipisicing elit.
                            Aperiam</p>
                    </Card.Content>
                </Card>
                <Card className="bg-white rounded-2xl p-10">
                    <Card.Header>
                        <h2 className="text-2xl font-bold">Free</h2>
                    </Card.Header>
                    <Card.Content>
                        <p className="text-gray-500">Lorem ipsum dolor sit amet, consectetur adipisicing elit.
                            Aperiam</p>
                    </Card.Content>
                </Card>
                <Card className="bg-white rounded-2xl p-10">
                    <Card.Header>
                        <h2 className="text-2xl font-bold">Free</h2>
                    </Card.Header>
                    <Card.Content>
                        <p className="text-gray-500">Lorem ipsum dolor sit amet, consectetur adipisicing elit.
                            Aperiam</p>
                    </Card.Content>
                </Card>
            </section>
        </section>
    )
}