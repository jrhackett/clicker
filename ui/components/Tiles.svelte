<script>
    import Tile from 'components/Tile'
    import colors from 'styles/colors'
    import { playerStore } from 'stores'

    let player
    const unsubscribe = playerStore.subscribe(p => player = p);

    var groupBy = function(xs, key) {
        return xs.reduce(function(rv, x) {
            (rv[x[key]] = rv[x[key]] || []).push(x);
            return rv;
        }, {});
    };

    $: categories = !!player ? player.generators.reduce(function(acc, g) {
            (acc[g.category] = acc[g.category] || []).push(g);
            return acc;
        }, {}) : {};
</script>

<style>
    .outer {
        padding: 1rem 2rem;
    }

    .inner {
        display: flex;
        flex-wrap: wrap;
    }

    li {
        list-style: none;
    }

    div {
        padding: 0 0.5rem;
    }

    p {
        color: var(--text-color);
        padding: 0.25rem 0;
        border-bottom: 1px solid var(--text-color);
    }
</style>

<ul class="outer" style="--text-color:{ colors.white }">
    {#each Object.keys(categories) as category}
        <li>
            <div>
                <p>{ category }</p>
            </div>
            <ul class="inner">
                {#each categories[category] as generator}
                    <Tile generator={ generator } />
                {/each}
            </ul>
        </li>
    {/each}
</ul>