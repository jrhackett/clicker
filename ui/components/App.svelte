<script>
	import Nav from 'components/nav'
	import colors from 'styles/colors'
    import router, { curRoute } from 'history/router'
    import RouterLink from 'components/nav/RouterLink'
    import { onMount } from 'svelte'

    onMount(() => {
        curRoute.set(window.location.pathname)
        if (!history.state) {
            window.history.replaceState({path: window.location.pathname}, '',   window.location.href)
        }
    })
    
    const handleBackNavigation = e => curRoute.set(e.state.path)

    let navHeight = '3rem'
    let backgroundColor = colors.offWhite
</script>

<style>
	.app {
        background-color: var(--background-color);
        height: 100vh;
    }

    .container {
        padding: 1rem 4rem;
    }
    
    #pageContent {
        padding-top: var(--nav-height);
    }

    #pageContent, .container {
        display: flex;
        flex-direction: column;
        flex: 1;
    }
</style>

<svelte:window on:popstate={ handleBackNavigation } />
<Nav navHeight={ navHeight } />
<div class="app" style="--background-color:{ backgroundColor };">
    <div id="pageContent" style="--nav-height:{ navHeight };">
        <div class="container">
            <svelte:component this={ router[$curRoute] } />
        </div>
    </div>
</div>