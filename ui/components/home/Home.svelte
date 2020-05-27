<script>
    export let header = 'Clicker'
</script>

<style>
    h1 {
        font-size: 2rem;
        font-weight: 600;
        margin-bottom: 0.5rem;
        flex-basis: 100%;
    }

    .container {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100%;
        flex-wrap: wrap;
    }
    .offset {
        margin-bottom: 20%;
    }
</style>

<div class="container">
    <div class="offset">
        <h1>{ header }</h1>
    </div>
</div>
