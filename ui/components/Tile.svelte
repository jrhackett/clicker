<script>
    import colors from 'styles/colors'
    import { updatePlayer, updateGeneratorsCost } from 'stores/updates'
    import { playerStore } from 'stores'

    export let generator = {}

    let player
    const unsubscribe = playerStore.subscribe(p => player = p);

    const handleClick = async name => await fetch('/api/v1/buy', {
        method: 'POST',
        body: JSON.stringify({ name })
    })
        .then(r => r.json())
        .then(async data => {
            updatePlayer(data.player)
            await updateGeneratorsCost(name)
        })

    $: buttonColor = player.liquid >= generator.cost ? colors.yellow : colors.lightGrey
    $: buttonHoverColor = player.liquid >= generator.cost ? colors.brightYellow : colors.lightGrey
    $: buttonSelectedColor = player.liquid >= generator.cost ? colors.dullYellow : colors.lightGrey
</script>

<style>
    li {
        list-style: none;
        flex-basis: 33.33%;
    }

    .inner {
        padding: 0.5rem;
    }

    .tile {
        display: flex;
        padding: 0.5rem 1rem;
        border-radius: 2px; 
        color: var(--text-color);
        background-color: var(--tile-color);
    }

    .col {
        display: flex;
        flex-direction: column;
        padding: 0 0.25rem;
    }

    h3 {
        font-size: 1rem;
    }

    p, span {
        font-size: 0.7rem;
    }

    span {
        font-weight: bold;
    }

    button {
        background-color: var(--button-color);
        border: none;
    }

    button:active, button:focus {
        outline: none;
    }

    button:hover {
        background-color: var(--button-hover-color);
    }

    button:active {
        background-color: var(--button-selected-color);
    }
</style>

<li>
    <div class="inner">
        <div class="tile" style="--tile-color:{ colors.blue };--text-color:{ colors.black }">
            <div class="col">
                <h3>{ generator.name }</h3>
                {#if !!generator.image_data_uri}
                    <img src={ generator.image_data_uri } alt={ generator.name } />
                {/if}
            </div>
            <div>
                <p>Cost: <span>{ generator.cost && generator.cost.toFixed(0) } likes</span></p>
                {#if generator.count > 0}
                    <p>Owned: <span>{ generator.count }</span></p>
                    <p>Earned: <span>{ generator.gained.toFixed(0) }</span></p>
                {/if}
                <p>Yield: <span>{ generator.gain_per_second }</span></p>
                <button
                    on:click={ () => handleClick(generator.name) }
                    style="--button-color:{ buttonColor };--button-hover-color:{ buttonHoverColor };--button-selected-color:{ buttonSelectedColor }"
                >
                    Buy
                </button>
            </div>
        </div>
    </div>
</li>
