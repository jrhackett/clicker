<script>
    import { playerStore } from 'stores'
    
    let player
    const unsubscribe = playerStore.subscribe(p => player = p);
</script>

<style>
    div {
        display: flex;
        align-items: center;
        flex-direction: column;
    }

    p {
        padding: 0.5rem 0;
    }
</style>

<div>
    <p>{ player.liquid.toFixed(0) } likes</p>
    <p>{ player.generators.reduce((acc, g) => acc += g.count * g.gain_per_second, 0) }lps</p>
</div>
