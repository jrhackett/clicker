<script>
    import { curRoute } from 'history/router'
    import colors from 'styles/colors'

    export let page = {
        path: '/',
        name: 'Home'
    }

    export let color = colors.white

    const redirectTo = e  => {
        curRoute.set(e.target.pathname)
        window.history.pushState({path: page.path}, '', window.location.origin + page.path)
    }
</script>

<style>
    a {
        display: flex;
        align-items: center;
        text-decoration: none;
        color: var(--text-color);
        height: 100%;
    }
</style>

<a href={ page.path } on:click|preventDefault={ redirectTo } style="--text-color:{ color };">
    { page.name }
</a>
