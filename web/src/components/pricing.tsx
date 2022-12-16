import {Button, Card} from "@/components";

export const Pricing = ({children, ...restProps}: any) => {
    return (
        <section {...restProps}>
            {children}
        </section>
    )
}

Pricing.Item = ({title, description, price, features, ...restProps}: any) => {
    return (
        <Card {...restProps}>
            <Card.Header>
                <h3 className="text-2xl font-bold">{title}</h3>
                <p className="text-xl">{description}</p>
            </Card.Header>
            <Card.Content>
                <ul className="list-disc list-inside">
                    {
                        features.map(({id, title, description}: any) => (
                            <li key={id} className="list-none my-10">
                                <h4 className="text-xl font-bold">{title}</h4>
                                <p>{description}</p>
                            </li>
                        ))
                    }

                </ul>
                <div>
                    <span className="text-3xl">{price}</span>
                </div>
                <div className="mt-5">
                    <Button className="w-full">Go Starter</Button>
                </div>
            </Card.Content>
        </Card>
    )
}