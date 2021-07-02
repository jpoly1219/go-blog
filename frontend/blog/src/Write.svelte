<script>
    function resize() {
        this.style.height = 'auto';
        this.style.height = this.scrollHeight + 'px'; 
    }
</script>

<div class="container mx-auto px-96 pb-10 flex flex-col">
    <div class="flex flex-col my-5">
        <input type="text" placeholder="Title" class="place-self-stretch border-none p-2 text-2xl font-medium mb-5 text-gray-900">
        <textarea on:input={resize} placeholder="Tell your story..." class="place-self-stretch border-none p-2 font-medium text-gray-900"></textarea>
    </div>
    <div class="flex flex-row place-content-center my-5">
        <button on:click={() => {alert("post submitted!")}} class="inline-flex items-center bg-blue-400 border rounded-lg text-base p-3 mr-5">
            <span class="text-white">Publish</span>
        </button>
        <button on:click={() => {alert("canceled")}} class="inline-flex items-center border border-gray-500 rounded-lg p-3">
            <span class="text-gray-900">Cancel</span>
        </button>
    </div>
</div>