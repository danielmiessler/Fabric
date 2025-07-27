<script lang="ts">
  import { X } from 'lucide-svelte';
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher<{
    close: void;
  }>();

  let modalContent: HTMLDivElement;

  function handleMouseLeave() {
    dispatch('close');
  }
</script>



<div class="bg-primary-800 rounded-lg flex flex-col h-[85vh] w-[600px] shadow-lg">
    <div class="flex justify-between items-center p-6 border-b-2 border-primary-700">
      <h1 class="text-3xl font-bold text-primary-300">Help</h1>
        <button class="text-muted-foreground hover:text-primary-300 transition-colors" on:click={() => dispatch('close')}>✕</button>
    </div>
    <div class="p-6 flex-1 overflow-y-auto" bind:this={modalContent}>
        <div class="space-y-4">
            <section>
                <h3 class="text-base font-bold text-primary-300 mb-2">Getting Started</h3>
                <p class="text-sm text-muted-foreground">Click on "Pattern Description" at the top to select a pattern or select a pattern directly from the dropdown menu.</p>
            </section>

            <section>
                <h3 class="text-base font-bold text-primary-300 mb-2">YouTube & other website URL</h3>
                <p class="text-sm text-muted-foreground">Paste a YouTube URL or an article URL in the message box and the link will be processed automatically. Youtube transcripts will be fetch and website html will be converted to markdown for LLM processing. Make sure to setup your Jina API key.</p>
            </section>

            <section>
                <h3 class="text-base font-bold text-primary-300 mb-2">Language Support</h3>
                <p class="text-sm text-muted-foreground">Select your preferred language from the dropdown menu. The AI will respond in the selected language.</p>
            </section>

            <section>
                <h3 class="text-base font-bold text-primary-300 mb-2">Model Configuration</h3>
                <p class="text-sm text-muted-foreground">Adjust model parameters like temperature and max length to control the AI's output. Higher temperature means more creative but potentially less focused responses.</p>
            </section>

            <section>
                <h3 class="text-base font-bold text-primary-300 mb-2">I only want a transcript or quick summary</h3>
                <p class="text-sm text-muted-foreground">Then don't select a pattern and simply ask in the message box "Give me a transcript" and paste the URL. it's that simple.</p>
            </section>

            <section>
                <h3 class="text-base font-bold text-primary-300 mb-2">Upload PDFs and other txt files</h3>
                <p class="text-sm text-muted-foreground">Now accepts PDFs and other txt files as input using the paperclip icon. PDFs are converted to markdown automatically.</p>
            </section>

            <section>
                <h3 class="text-base font-bold text-primary-300 mb-2">Taking notes and saving to Obsidian</h3>
                <p class="text-sm text-muted-foreground">By default, your personal folders should be located exactly here: web/myfiles/Fabric_obsidian and web/myfiles/inbox for notes.  </p>
            </section>

            <section>
              <h3 class="text-base font-bold text-primary-300 mb-2">Managing Pattern Descriptions and Tags</h3>
              <p class="text-sm text-muted-foreground">Refer to instructions in the WEB INTERFACE MOD README FILES folder.  </p>
          </section>

            <section class="mt-12 pt-8 border-t-2 border-primary-700">
              <h1 class="text-2xl font-bold text-primary-300 mb-4">PATTERN TAGS</h1>
              <p class="text-sm text-muted-foreground mb-6">You can configure the TAGs as you wish and modify or replace these TAGs with yours.</p>
              <div class="text-sm text-muted-foreground space-y-4">
                    <div>
                      <p class="text-base font-bold text-primary-300 mb-2">SECURITY</p>
                        <p>Any pattern pertaining to IT Security</p>
                    </div>

                    <div>
                      <p class="text-base font-bold text-primary-300 mb-2">WRITING</p>
                        <p>Any pattern pertaining to writing, editing, documentation, communication</p>
                    </div>

                    <div>
                      <p class="text-base font-bold text-primary-300 mb-2">ANALYSIS</p>
                        <p>For patterns like analyze_paper, analyze_claims, analyze_debate</p>
                    </div>

                    <div>
                      <p class="text-base font-bold text-primary-300 mb-2">VISUALIZATION</p>
                        <p>Any pattern involving some visual representation of something: For example, create logo, presentation diagrams etc.</p>
                    </div>

                    <div>
                      <p class="text-base font-bold text-primary-300 mb-2">DEVELOPMENT</p>
                        <p>Any pattern related to software development. For example, patterns like create_coding_project, create_prd, write_pull-request. Covers software development, coding, technical documentation</p>
                    </div>

                    <div>
                      <p class="text-base font-bold text-primary-300 mb-2">SUMMARIZATION</p>
                        <p>Covers content condensing and key point extraction. For patterns like create_5_sentence_summary, summarize_meeting, create_micro_summary</p>
                    </div>

                    <div>
                      <p class="text-base font-bold text-primary-300 mb-2">EXTRACTION</p>
                        <p>For patterns where key focus is mining specific information from content. For example, extract_wisdom, extract_skills, extract_patterns.</p>
                    </div>

                    <div>
                      <p class="text-base font-bold text-primary-300 mb-2">STRATEGY</p>
                        <p>For patterns like analyze_military_strategy, prepare_7s_strategy. Covers planning, decision-making frameworks, mostly with business focus.</p>
                    </div>

                    <div>
                      <p class="text-base font-bold text-primary-300 mb-2">LEARNING</p>
                        <p>Covers educational content and teaching. For patterns like to_flashcards, create_quiz, explain_math</p>
                    </div>

                    <div>
                      <p class="text-base font-bold text-primary-300 mb-2">CONVERSION</p>
                        <p>Covers format and language transformation. For patterns like convert_to_markdown, translate, sanitize_broken_html_to_markdown</p>
                    </div>

                    <div>
                      <p class="text-base font-bold text-primary-300 mb-2">REVIEW</p>
                        <p>Covers evaluation and assessment of source. For patterns like review_design, rate_content, analyze_presentation</p>
                    </div>

                    <div>
                      <p class="text-base font-bold text-primary-300 mb-2">BUSINESS</p>
                        <p>Pattern aiming to support a business or entrepreneurial aim. For patterns like create_hormozi_offer, extract_business_ideas</p>
                    </div>

                    <div>
                      <p class="text-base font-bold text-primary-300 mb-2">AI</p>
                        <p>For patterns like improve_prompt, rate_ai_response. Covers AI-specific interactions</p>
                    </div>

                    <div>
                      <p class="text-base font-bold text-primary-300 mb-2">GAMING</p>
                        <p>For patterns like create_npc, create_rpg_summary. Covers gaming and simulation content</p>
                    </div>

                    <div>
                      <p class="text-base font-bold text-primary-300 mb-2">RESEARCH</p>
                        <p>For patterns like analyze_paper, create_academic_paper. Covers academic and scientific content</p>
                    </div>

                    <div>
                      <p class="text-base font-bold text-primary-300 mb-2">CRITICAL THINKING</p>
                        <p>Any pattern aiming to improve someone's thinking or aiming to assess evidence.</p>
                    </div>

                    <div>
                      <p class="text-base font-bold text-primary-300 mb-2">SELF</p>
                        <p>Any pattern dealing with personal growth or clearly focus on personal outcome.</p>
                    </div>

                    <div>
                      <p class="text-base font-bold text-primary-300 mb-2">WISDOM</p>
                        <p>Anything worthy of adding to your personal playbook.</p>
                    </div>

                    <div>
                      <p class="text-base font-bold text-primary-300 mb-2">OTHER</p>
                        <p>Does not fit anywhere</p>
                    </div>
                </div>
            </section>
        </div>
    </div>
</div>

