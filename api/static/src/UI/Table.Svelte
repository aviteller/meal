<script>
  export let headers;
  export let title = null
</script>

<style>
  table {
    width: 100%;
    text-align: left;
  }
</style>

{#if title}
  <h1>{title}</h1>
{/if}

<table>
  <thead>
    <tr>
      {#each headers as header}
        <th>{header}</th>
      {/each}
    </tr>
  </thead>
  <tbody>
    <slot />
  </tbody>
</table>
