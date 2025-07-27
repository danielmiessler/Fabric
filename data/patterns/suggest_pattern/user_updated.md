# Suggest Pattern

## OVERVIEW

What It Does: Fabric is an open-source framework designed to augment human capabilities using AI, making it easier to integrate AI into daily tasks.

Why People Use It: Users leverage Fabric to seamlessly apply AI for solving everyday challenges, enhancing productivity, and fostering human creativity through technology.

## HOW TO USE IT

Most Common Syntax: The most common usage involves executing Fabric commands in the terminal, such as `fabric --pattern <PATTERN_NAME>`.

## COMMON USE CASES

For Summarizing Content: `fabric --pattern summarize`
For Analyzing Claims: `fabric --pattern analyze_claims`
For Extracting Wisdom from Videos: `fabric --pattern extract_wisdom`
For creating custom patterns: `fabric --pattern create_pattern`

- One possible place to store them is ~/.config/custom-fabric-patterns.
- Then when you want to use them, simply copy them into ~/.config/fabric/patterns.
`cp -a ~/.config/custom-fabric-patterns/* ~/.config/fabric/patterns/`
- Now you can run them with: `pbpaste | fabric -p your_custom_pattern`

## MOST IMPORTANT AND USED OPTIONS AND FEATURES

- **--pattern PATTERN, -p PATTERN**: Specifies the pattern (prompt) to use. Useful for applying specific AI prompts to your input.

- **--stream, -s**: Streams results in real-time. Ideal for getting immediate feedback from AI operations.

- **--update, -u**: Updates patterns. Ensures you're using the latest AI prompts for your tasks.

- **--model MODEL, -m MODEL**: Selects the AI model to use. Allows customization of the AI backend for different tasks.

- **--setup, -S**: Sets up your Fabric instance. Essential for first-time users to configure Fabric correctly.

- **--list, -l**: Lists available patterns. Helps users discover new AI prompts for various applications.

- **--context, -C**: Uses a Context file to add context to your pattern. Enhances the relevance of AI responses by providing additional background information.

## PATTERNS

**Key pattern to use: `suggest_pattern`** - suggests appropriate fabric patterns or commands based on user input.

### agility_story

Generate a user story and acceptance criteria in JSON format based on the given topic.

### ai

Interpret questions deeply and provide concise, insightful answers in Markdown bullet points.

### analyze_answers

Evaluate quiz answers for correctness based on learning objectives and generated quiz questions.

### analyze_bill

Analyzes legislation to identify overt and covert goals, examining bills for hidden agendas and true intentions.

### analyze_bill_short

Provides a concise analysis of legislation, identifying overt and covert goals in a brief, structured format.

### analyze_candidates

Compare and contrast two political candidates based on key issues and policies.

### analyze_cfp_submission

Review and evaluate conference speaking session submissions based on clarity, relevance, depth, and engagement potential.

### analyze_claims

Analyse and rate truth claims with evidence, counter-arguments, fallacies, and final recommendations.

### analyze_comments

Evaluate internet comments for content, categorize sentiment, and identify reasons for praise, criticism, and neutrality.

### analyze_debate

Rate debates on insight, emotionality, and present an unbiased, thorough analysis of arguments, agreements, and disagreements.

### analyze_email_headers

Provide cybersecurity analysis and actionable insights on SPF, DKIM, DMARC, and ARC email header results.

### analyze_incident

Efficiently extract and organize key details from cybersecurity breach articles, focusing on attack type, vulnerable components, attacker and target info, incident details, and remediation steps.

### analyze_interviewer_techniques

This exercise involves analyzing interviewer techniques, identifying their unique qualities, and succinctly articulating what makes them stand out in a clear, simple format.

### analyze_logs

Analyse server log files to identify patterns, anomalies, and issues, providing data-driven insights and recommendations for improving server reliability and performance.

### analyze_malware

Analyse malware details, extract key indicators, techniques, and potential detection strategies, and summarize findings concisely for a malware analyst's use in identifying and responding to threats.

### analyze_military_strategy

Analyse a historical battle, offering in-depth insights into strategic decisions, strengths, weaknesses, tactical approaches, logistical factors, pivotal moments, and consequences for a comprehensive military evaluation.

### analyze_mistakes

Analyse past mistakes in thinking patterns, map them to current beliefs, and offer recommendations to improve accuracy in predictions.

### analyze_paper

Analyses research papers by summarizing findings, evaluating rigor, and assessing quality to provide insights for documentation and review.

### analyze_paper_simple

Analyzes academic papers with a focus on primary findings, research quality, and study design evaluation.

### analyze_patent

Analyse a patent's field, problem, solution, novelty, inventive step, and advantages in detail while summarizing and extracting keywords.

### analyze_personality

Performs a deep psychological analysis of a person in the input, focusing on their behavior, language, and psychological traits.

### analyze_presentation

Reviews and critiques presentations by analyzing the content, speaker's underlying goals, self-focus, and entertainment value.

### analyze_product_feedback

A prompt for analyzing and organizing user feedback by identifying themes, consolidating similar comments, and prioritizing them based on usefulness.

### analyze_proposition

Analyzes a ballot proposition by identifying its purpose, impact, arguments for and against, and relevant background information.

### analyze_prose

Evaluates writing for novelty, clarity, and prose, providing ratings, improvement recommendations, and an overall score.

### analyze_prose_json

Evaluates writing for novelty, clarity, prose, and provides ratings, explanations, improvement suggestions, and an overall score in a JSON format.

### analyze_prose_pinker

Evaluates prose based on Steven Pinker's The Sense of Style, analyzing writing style, clarity, and bad writing elements.

### analyze_risk

Conducts a risk assessment of a third-party vendor, assigning a risk score and suggesting security controls based on analysis of provided documents and vendor website.

### analyze_sales_call

Rates sales call performance across multiple dimensions, providing scores and actionable feedback based on transcript analysis.

### analyze_spiritual_text

Compares and contrasts spiritual texts by analyzing claims and differences with the King James Bible.

### analyze_tech_impact

Analyzes the societal impact, ethical considerations, and sustainability of technology projects, evaluating their outcomes and benefits.

### analyze_terraform_plan

Analyzes Terraform plan outputs to assess infrastructure changes, security risks, cost implications, and compliance considerations.

### analyze_threat_report

Extracts surprising insights, trends, statistics, quotes, references, and recommendations from cybersecurity threat reports, summarizing key findings and providing actionable information.

### analyze_threat_report_cmds

Extract and synthesize actionable cybersecurity commands from provided materials, incorporating command-line arguments and expert insights for pentesters and non-experts.

### analyze_threat_report_trends

Extract up to 50 surprising, insightful, and interesting trends from a cybersecurity threat report in markdown format.

### answer_interview_question

Generates concise, tailored responses to technical interview questions, incorporating alternative approaches and evidence to demonstrate the candidate's expertise and experience.

### ask_secure_by_design_questions

Generates a set of security-focused questions to ensure a project is built securely by design, covering key components and considerations.

### ask_uncle_duke

Coordinates a team of AI agents to research and produce multiple software development solutions based on provided specifications, and conducts detailed code reviews to ensure adherence to best practices.

### capture_thinkers_work

Analyze philosophers or philosophies and provide detailed summaries about their teachings, background, works, advice, and related concepts in a structured template.

### check_agreement

Analyze contracts and agreements to identify important stipulations, issues, and potential gotchas, then summarize them in Markdown.

### clean_text

Fix broken or malformatted text by correcting line breaks, punctuation, capitalization, and paragraphs without altering content or spelling.

### coding_master

Explain a coding concept to a beginner, providing examples, and formatting code in markdown with specific output sections like ideas, recommendations, facts, and insights.

### compare_and_contrast

Compare and contrast a list of items in a markdown table, with items on the left and topics on top.

### convert_to_markdown

Convert content to clean, complete Markdown format, preserving all original structure, formatting, links, and code blocks without alterations.

### create_5_sentence_summary

Create concise summaries or answers to input at 5 different levels of depth, from 5 words to 1 word.

### create_academic_paper

Generate a high-quality academic paper in LaTeX format with clear concepts, structured content, and a professional layout.

### create_ai_jobs_analysis

Analyze job categories' susceptibility to automation, identify resilient roles, and provide strategies for personal adaptation to AI-driven changes in the workforce.

### create_aphorisms

Find and generate a list of brief, witty statements.

### create_art_prompt

Generates a detailed, compelling visual description of a concept, including stylistic references and direct AI instructions for creating art.

### create_better_frame

Identifies and analyzes different frames of interpreting reality, emphasizing the power of positive, productive lenses in shaping outcomes.

### create_coding_feature

Generates secure and composable code features using modern technology and best practices from project specifications.

### create_coding_project

Generate wireframes and starter code for any coding ideas that you have.

### create_command

Helps determine the correct parameters and switches for penetration testing tools based on a brief description of the objective.

### create_cyber_summary

Summarizes cybersecurity threats, vulnerabilities, incidents, and malware with a 25-word summary and categorized bullet points, after thoroughly analyzing and mapping the provided input.

### create_design_document

Creates a detailed design document for a system using the C4 model, addressing business and security postures, and including a system context diagram.

### create_diy

Creates structured "Do It Yourself" tutorial patterns by analyzing prompts, organizing requirements, and providing step-by-step instructions in Markdown format.

### create_excalidraw_visualization

Creates complex Excalidraw diagrams to visualize relationships between concepts and ideas in structured format.

### create_flash_cards

Creates flashcards for key concepts, definitions, and terms with question-answer format for educational purposes.

### create_formal_email

Crafts professional, clear, and respectful emails by analyzing context, tone, and purpose, ensuring proper structure and formatting.

### create_git_diff_commit

Generates Git commands and commit messages for reflecting changes in a repository, using conventional commits and providing concise shell commands for updates.

### create_graph_from_input

Generates a CSV file with progress-over-time data for a security program, focusing on relevant metrics and KPIs.

### create_hormozi_offer

Creates a customized business offer based on principles from Alex Hormozi's book, "$100M Offers."

### create_idea_compass

Organizes and structures ideas by exploring their definition, evidence, sources, and related themes or consequences.

### create_investigation_visualization

Creates detailed Graphviz visualizations of complex input, highlighting key aspects and providing clear, well-annotated diagrams for investigative analysis and conclusions.

### create_keynote

Creates TED-style keynote presentations with a clear narrative, structured slides, and speaker notes, emphasizing impactful takeaways and cohesive flow.

### create_loe_document

Creates detailed Level of Effort documents for estimating work effort, resources, and costs for tasks or projects.

### create_logo

Creates simple, minimalist company logos without text, generating AI prompts for vector graphic logos based on input.

### create_markmap_visualization

Transforms complex ideas into clear visualizations using MarkMap syntax, simplifying concepts into diagrams with relationships, boxes, arrows, and labels.

### create_mermaid_visualization

Creates detailed, standalone visualizations of concepts using Mermaid (Markdown) syntax, ensuring clarity and coherence in diagrams.

### create_mermaid_visualization_for_github

Creates standalone, detailed visualizations using Mermaid (Markdown) syntax to effectively explain complex concepts, ensuring clarity and precision.

### create_micro_summary

Summarizes content into a concise, 20-word summary with main points and takeaways, formatted in Markdown.

### create_mnemonic_phrases

Creates memorable mnemonic sentences from given words to aid in memory retention and learning.

### create_network_threat_landscape

Analyzes open ports and services from a network scan and generates a comprehensive, insightful, and detailed security threat report in Markdown.

### create_newsletter_entry

Condenses provided article text into a concise, objective, newsletter-style summary with a title in the style of Frontend Weekly.

### create_npc

Generates a detailed D&D 5E NPC, including background, flaws, stats, appearance, personality, goals, and more in Markdown format.

### create_pattern

Extracts, organizes, and formats LLM/AI prompts into structured sections, detailing the AI's role, instructions, output format, and any provided examples for clarity and accuracy.

### create_prd

Creates a precise Product Requirements Document (PRD) in Markdown based on input.

### create_prediction_block

Extracts and formats predictions from input into a structured Markdown block for a blog post.

### create_quiz

Generates review questions based on learning objectives from the input, adapted to the specified student level, and outputs them in a clear markdown format.

### create_reading_plan

Creates a three-phase reading plan based on an author or topic to help the user become significantly knowledgeable, including core, extended, and supplementary readings.

### create_recursive_outline

Breaks down complex tasks or projects into manageable, hierarchical components with recursive outlining for clarity and simplicity.

### create_report_finding

Creates a detailed, structured security finding report in markdown, including sections on Description, Risk, Recommendations, References, One-Sentence-Summary, and Quotes.

### create_rpg_summary

Summarizes an in-person RPG session with key events, combat details, player stats, and role-playing highlights in a structured format.

### create_security_update

Creates concise security updates for newsletters, covering stories, threats, advisories, vulnerabilities, and a summary of key issues.

### create_show_intro

Creates compelling short intros for podcasts, summarizing key topics and themes discussed in the episode.

### create_sigma_rules

Extracts Tactics, Techniques, and Procedures (TTPs) from security news and converts them into Sigma detection rules for host-based detections.

### create_story_explanation

Summarizes complex content in a clear, approachable story format that makes the concepts easy to understand.

### create_stride_threat_model

Create a STRIDE-based threat model for a system design, identifying assets, trust boundaries, data flows, and prioritizing threats with mitigations.

### create_summary

Summarizes content into a 20-word sentence, 10 main points (16 words max), and 5 key takeaways in Markdown format.

### create_tags

Identifies at least 5 tags from text content for mind mapping tools, including authors and existing tags if present.

### create_threat_scenarios

Identifies likely attack methods for any system by providing a narrative-based threat model, balancing risk and opportunity.

### create_ttrc_graph

Creates a CSV file showing the progress of Time to Remediate Critical Vulnerabilities over time using given data.

### create_ttrc_narrative

Creates a persuasive narrative highlighting progress in reducing the Time to Remediate Critical Vulnerabilities metric over time.

### create_upgrade_pack

Extracts world model and task algorithm updates from content, providing beliefs about how the world works and task performance.

### create_user_story

Writes concise and clear technical user stories for new features in complex software programs, formatted for all stakeholders.

### create_video_chapters

Extracts interesting topics and timestamps from a transcript, providing concise summaries of key moments.

### create_visualization

Transforms complex ideas into visualizations using intricate ASCII art, simplifying concepts where necessary.

### dialog_with_socrates

Engages in deep, meaningful dialogues to explore and challenge beliefs using the Socratic method.

### enrich_blog_post

Enhances Markdown blog files by applying instructions to improve structure, visuals, and readability for HTML rendering.

### explain_code

Explains code, security tool output, configuration text, and answers questions based on the provided input.

### explain_docs

Improves and restructures tool documentation into clear, concise instructions, including overviews, usage, use cases, and key features.

### explain_math

Helps you understand mathematical concepts in a clear and engaging way.

### explain_project

Summarizes project documentation into clear, concise sections covering the project, problem, solution, installation, usage, and examples.

### explain_terms

Produces a glossary of advanced terms from content, providing a definition, analogy, and explanation of why each term matters.

### export_data_as_csv

Extracts and outputs all data structures from the input in properly formatted CSV data.

### extract_algorithm_update_recommendations

Extracts concise, practical algorithm update recommendations from the input and outputs them in a bulleted list.

### extract_article_wisdom

Extracts surprising, insightful, and interesting information from content, categorizing it into sections like summary, ideas, quotes, facts, references, and recommendations.

### extract_book_ideas

Extracts and outputs 50 to 100 of the most surprising, insightful, and interesting ideas from a book's content.

### extract_book_recommendations

Extracts and outputs 50 to 100 practical, actionable recommendations from a book's content.

### extract_business_ideas

Extracts top business ideas from content and elaborates on the best 10 with unique differentiators.

### extract_controversial_ideas

Extracts and outputs controversial statements and supporting quotes from the input in a structured Markdown list.

### extract_core_message

Extracts and outputs a clear, concise sentence that articulates the core message of a given text or body of work.

### extract_ctf_writeup

Extracts a short writeup from a warstory-like text about a cyber security engagement.

### extract_domains

Extracts domains and URLs from content to identify sources used for articles, newsletters, and other publications.

### extract_extraordinary_claims

Extracts and outputs a list of extraordinary claims from conversations, focusing on scientifically disputed or false statements.

### extract_ideas

Extracts and outputs all the key ideas from input, presented as 15-word bullet points in Markdown.

### extract_insights

Extracts and outputs the most powerful and insightful ideas from text, formatted as 16-word bullet points in the INSIGHTS section, also IDEAS section.

### extract_insights_dm

Extracts and outputs all valuable insights and a concise summary of the content, including key points and topics discussed.

### extract_instructions

Extracts clear, actionable step-by-step instructions and main objectives from instructional video transcripts, organizing them into a concise list.

### extract_jokes

Extracts jokes from text content, presenting each joke with its punchline in separate bullet points.

### extract_latest_video

Extracts the latest video URL from a YouTube RSS feed and outputs the URL only.

### extract_main_activities

Extracts key events and activities from transcripts or logs, providing a summary of what happened.

### extract_main_idea

Extracts the main idea and key recommendation from the input, summarizing them in 15-word sentences.

### extract_most_redeeming_thing

Extracts the most redeeming aspect from an input, summarizing it in a single 15-word sentence.

### extract_patterns

Extracts and analyzes recurring, surprising, and insightful patterns from input, providing detailed analysis and advice for builders.

### extract_poc

Extracts proof of concept URLs and validation methods from security reports, providing the URL and command to run.

### extract_predictions

Extracts predictions from input, including specific details such as date, confidence level, and verification method.

### extract_primary_problem

Extracts the primary problem with the world as presented in a given text or body of work.

### extract_primary_solution

Extracts the primary solution for the world as presented in a given text or body of work.

### extract_product_features

Extracts and outputs a list of product features from the provided input in a bulleted format.

### extract_questions

Extracts and outputs all questions asked by the interviewer in a conversation or interview.

### extract_recipe

Extracts and outputs a recipe with a short meal description, ingredients with measurements, and preparation steps.

### extract_recommendations

Extracts and outputs concise, practical recommendations from a given piece of content in a bulleted list.

### extract_references

Extracts and outputs a bulleted list of references to art, stories, books, literature, and other sources from content.

### extract_skills

Extracts and classifies skills from a job description into a table, separating each skill and classifying it as either hard or soft.

### extract_song_meaning

Analyzes a song to provide a summary of its meaning, supported by detailed evidence from lyrics, artist commentary, and fan analysis.

### extract_sponsors

Extracts and lists official sponsors and potential sponsors from a provided transcript.

### extract_videoid

Extracts and outputs the video ID from any given URL.

### extract_wisdom

Extracts surprising, insightful, and interesting information from text on topics like human flourishing, AI, learning, and more.

### extract_wisdom_agents

Extracts valuable insights, ideas, quotes, and references from content, emphasizing topics like human flourishing, AI, learning, and technology.

### extract_wisdom_dm

Extracts all valuable, insightful, and thought-provoking information from content, focusing on topics like human flourishing, AI, learning, and technology.

### extract_wisdom_nometa

Extracts insights, ideas, quotes, habits, facts, references, and recommendations from content, focusing on human flourishing, AI, technology, and related topics.

### find_female_life_partner

Analyzes criteria for finding a female life partner and provides clear, direct, and poetic descriptions.

### find_hidden_message

Extracts overt and hidden political messages, justifications, audience actions, and a cynical analysis from content.

### find_logical_fallacies

Identifies and analyzes fallacies in arguments, classifying them as formal or informal with detailed reasoning.

### get_wow_per_minute

Determines the wow-factor of content per minute based on surprise, novelty, insight, value, and wisdom, measuring how rewarding the content is for the viewer.

### get_youtube_rss

Returns the RSS URL for a given YouTube channel based on the channel ID or URL.

### humanize

Rewrites AI-generated text to sound natural, conversational, and easy to understand, maintaining clarity and simplicity.

### identify_dsrp_distinctions

Encourages creative, systems-based thinking by exploring distinctions, boundaries, and their implications, drawing on insights from prominent systems thinkers.

### identify_dsrp_perspectives

Explores the concept of distinctions in systems thinking, focusing on how boundaries define ideas, influence understanding, and reveal or obscure insights.

### identify_dsrp_relationships

Encourages exploration of connections, distinctions, and boundaries between ideas, inspired by systems thinkers to reveal new insights and patterns in complex systems.

### identify_dsrp_systems

Encourages organizing ideas into systems of parts and wholes, inspired by systems thinkers to explore relationships and how changes in organization impact meaning and understanding.

### identify_job_stories

Identifies key job stories or requirements for roles.

### improve_academic_writing

Refines text into clear, concise academic language while improving grammar, coherence, and clarity, with a list of changes.

### improve_prompt

Improves an LLM/AI prompt by applying expert prompt writing strategies for better results and clarity.

### improve_report_finding

Improves a penetration test security finding by providing detailed descriptions, risks, recommendations, references, quotes, and a concise summary in markdown format.

### improve_writing

Refines text by correcting grammar, enhancing style, improving clarity, and maintaining the original meaning.

### judge_output

Evaluates Honeycomb queries by judging their effectiveness, providing critiques and outcomes based on language nuances and analytics relevance.

### label_and_rate

Labels content with up to 20 single-word tags and rates it based on idea count and relevance to human meaning, AI, and other related themes, assigning a tier (S, A, B, C, D) and a quality score.

### md_callout

Classifies content and generates a markdown callout based on the provided text, selecting the most appropriate type.

### official_pattern_template

Template to use if you want to create new fabric patterns.

### prepare_7s_strategy

Prepares a comprehensive briefing document from 7S's strategy capturing organizational profile, strategic elements, and market dynamics with clear, concise, and organized content.

### provide_guidance

Provides psychological and life coaching advice, including analysis, recommendations, and potential diagnoses, with a compassionate and honest tone.

### rate_ai_response

Rates the quality of AI responses by comparing them to top human expert performance, assigning a letter grade, reasoning, and providing a 1-100 score based on the evaluation.

### rate_ai_result

Assesses the quality of AI/ML/LLM work by deeply analyzing content, instructions, and output, then rates performance based on multiple dimensions, including coverage, creativity, and interdisciplinary thinking.

### rate_content

Labels content with up to 20 single-word tags and rates it based on idea count and relevance to human meaning, AI, and other related themes, assigning a tier (S, A, B, C, D) and a quality score.

### rate_value

Produces the best possible output by deeply analyzing and understanding the input and its intended purpose.

### raw_query

Fully digests and contemplates the input to produce the best possible result based on understanding the sender's intent.

### recommend_artists

Recommends a personalized festival schedule with artists aligned to your favorite styles and interests, including rationale.

### recommend_pipeline_upgrades

Optimizes vulnerability-checking pipelines by incorporating new information and improving their efficiency, with detailed explanations of changes.

### recommend_talkpanel_topics

Produces a clean set of proposed talks or panel talking points for a person based on their interests and goals, formatted for submission to a conference organizer.

### refine_design_document

Refines a design document based on a design review by analyzing, mapping concepts, and implementing changes using valid Markdown.

### review_design

Reviews and analyzes architecture design, focusing on clarity, component design, system integrations, security, performance, scalability, and data management.

### sanitize_broken_html_to_markdown

Converts messy HTML into clean, properly formatted Markdown, applying custom styling and ensuring compatibility with Vite.

### show_fabric_options_markmap

Visualizes the functionality of the Fabric framework by representing its components, commands, and features based on the provided input.

### solve_with_cot

Provides detailed, step-by-step responses with chain of thought reasoning, using structured thinking, reflection, and output sections.

### suggest_pattern

Suggests appropriate fabric patterns or commands based on user input, providing clear explanations and options for users.

### summarize

Summarizes content into a 20-word sentence, main points, and takeaways, formatted with numbered lists in Markdown.

### summarize_board_meeting

Creates formal meeting notes from board meeting transcripts for corporate governance documentation.

### summarize_debate

Summarizes debates, identifies primary disagreement, extracts arguments, and provides analysis of evidence and argument strength to predict outcomes.

### summarize_git_changes

Summarizes recent project updates from the last 7 days, focusing on key changes with enthusiasm.

### summarize_git_diff

Summarizes and organizes Git diff changes with clear, succinct commit messages and bullet points.

### summarize_lecture

Extracts relevant topics, definitions, and tools from lecture transcripts, providing structured summaries with timestamps and key takeaways.

### summarize_legislation

Summarizes complex political proposals and legislation by analyzing key points, proposed changes, and providing balanced, positive, and cynical characterizations.

### summarize_meeting

Analyzes meeting transcripts to extract a structured summary, including an overview, key points, tasks, decisions, challenges, timeline, references, and next steps.

### summarize_micro

Summarizes content into a 20-word sentence, 3 main points, and 3 takeaways, formatted in clear, concise Markdown.

### summarize_newsletter

Extracts the most meaningful, interesting, and useful content from a newsletter, summarizing key sections such as content, opinions, tools, companies, and follow-up items in clear, structured Markdown.

### summarize_paper

Summarizes an academic paper by detailing its title, authors, technical approach, distinctive features, experimental setup, results, advantages, limitations, and conclusion in a clear, structured format using human-readable Markdown.

### summarize_prompt

Summarizes AI chat prompts by describing the primary function, unique approach, and expected output in a concise paragraph. The summary is focused on the prompt's purpose without unnecessary details or formatting.

### summarize_pull-requests

Summarizes pull requests for a coding project by providing a summary and listing the top PRs with human-readable descriptions.

### summarize_rpg_session

Summarizes a role-playing game session by extracting key events, combat stats, character changes, quotes, and more.

### t_analyze_challenge_handling

Provides 8-16 word bullet points evaluating how well challenges are being addressed, calling out any lack of effort.

### t_check_metrics

Analyzes deep context from the TELOS file and input instruction, then provides a wisdom-based output while considering metrics and KPIs to assess recent improvements.

### t_create_h3_career

Summarizes context and produces wisdom-based output by deeply analyzing both the TELOS File and the input instruction, considering the relationship between the two.

### t_create_opening_sentences

Describes from TELOS file the person's identity, goals, and actions in 4 concise, 32-word bullet points, humbly.

### t_describe_life_outlook

Describes from TELOS file a person's life outlook in 5 concise, 16-word bullet points.

### t_extract_intro_sentences

Summarizes from TELOS file a person's identity, work, and current projects in 5 concise and grounded bullet points.

### t_extract_panel_topics

Creates 5 panel ideas with titles and descriptions based on deep context from a TELOS file and input.

### t_find_blindspots

Identify potential blindspots in thinking, frames, or models that may expose the individual to error or risk.

### t_find_negative_thinking

Analyze a TELOS file and input to identify negative thinking in documents or journals, followed by tough love encouragement.

### t_find_neglected_goals

Analyze a TELOS file and input instructions to identify goals or projects that have not been worked on recently.

### t_give_encouragement

Analyze a TELOS file and input instructions to evaluate progress, provide encouragement, and offer recommendations for continued effort.

### t_red_team_thinking

Analyze a TELOS file and input instructions to red-team thinking, models, and frames, then provide recommendations for improvement.

### t_threat_model_plans

Analyze a TELOS file and input instructions to create threat models for a life plan and recommend improvements.

### t_visualize_mission_goals_projects

Analyze a TELOS file and input instructions to create an ASCII art diagram illustrating the relationship of missions, goals, and projects.

### t_year_in_review

Analyze a TELOS file to create insights about a person or entity, then summarize accomplishments and visualizations in bullet points.

### to_flashcards

Create Anki flashcards from a given text, focusing on concise, optimized questions and answers without external context.

### transcribe_minutes

Extracts (from meeting transcription) meeting minutes, identifying actionables, insightful ideas, decisions, challenges, and next steps in a structured format.

### translate

Translates sentences or documentation into the specified language code while maintaining the original formatting and tone.

### tweet

Provides a step-by-step guide on crafting engaging tweets with emojis, covering Twitter basics, account creation, features, and audience targeting.

### write_essay

Writes essays in the style of a specified author, embodying their unique voice, vocabulary, and approach. Uses `author_name` variable.

### write_essay_pg

Writes concise, clear essays in the style of Paul Graham, focusing on simplicity, clarity, and illumination of the provided topic.

### write_hackerone_report

Generates concise, clear, and reproducible bug bounty reports, detailing vulnerability impact, steps to reproduce, and exploit details for triagers.

### write_latex

Generates syntactically correct LaTeX code for a new.tex document, ensuring proper formatting and compatibility with pdflatex.

### write_micro_essay

Writes concise, clear, and illuminating essays on the given topic in the style of Paul Graham.

### write_nuclei_template_rule

Generates Nuclei YAML templates for detecting vulnerabilities using HTTP requests, matchers, extractors, and dynamic data extraction.

### write_pull-request

Drafts detailed pull request descriptions, explaining changes, providing reasoning, and identifying potential bugs from the git diff command output.

### write_semgrep_rule

Creates accurate and working Semgrep rules based on input, following syntax guidelines and specific language considerations.

### youtube_summary

Create concise, timestamped Youtube video summaries that highlight key points.
