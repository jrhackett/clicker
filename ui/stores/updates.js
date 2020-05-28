import { playerStore, userStore } from 'stores'

export const updatePlayer = newPlayer => {
    playerStore.update(p => ({
        ...p,
        liquid: newPlayer.liquid ? newPlayer.liquid : p.liquid,
        generators: newPlayer.generators.map(g => ({
            ...(p ? p.generators.find(a => a.name == g.name) : {}),
            ...g
        }))
    }))
}

export const updateUser = (id, name, email, image_url) => userStore.update(u => ({
    ...u,
    id,
    name,
    email,
    image_url
}))

export const updateGeneratorsCost = async name => {
    await fetch(`api/v1/cost?name=${ name }&count=1`)
        .then(r => r.json())
        .then(data => {
            playerStore.update(p => ({
                ...p,
                generators: p.generators.map(g => {
                    if (g.name == name) {
                        return {
                            ...g,
                            cost: data.cost
                        }
                    }

                    return g
                })
            }))
        })
}
