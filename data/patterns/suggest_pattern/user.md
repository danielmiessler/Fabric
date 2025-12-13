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

## PATTERNS BY CATEGORY

**Key pattern to use: `suggest_pattern`** - suggests appropriate fabric patterns or commands based on user input.

## AI PATTERNS

### ai

Provide concise, insightful answers in brief bullets focused on core concepts.

### create_art_prompt

Transform concepts into detailed AI art prompts with style references.

### create_pattern

Design structured patterns for AI prompts with identity, purpose, steps, output.

### create_prediction_block

Format predictions for tracking/verification in markdown prediction logs.

### extract_wisdom_agents

Extract insights from AI agent interactions, focusing on learning.

### improve_prompt

Enhance AI prompts by refining clarity and specificity.

### judge_output

Evaluate AI outputs for quality and accuracy.

### rate_ai_response

Evaluate AI responses for quality and effectiveness.

### rate_ai_result

Assess AI outputs against criteria, providing scores and feedback.

### raw_query

Process direct queries by interpreting intent.

### suggest_pattern

Recommend Fabric patterns based on user requirements.

## ANALYSIS PATTERNS

### analyze_answers

Evaluate student responses providing detailed feedback adapted to levels.

### analyze_bill

Analyze a legislative bill and implications.

### analyze_bill_short

Condensed - Analyze a legislative bill and implications.

### analyze_candidates

Compare candidate positions, policy differences and backgrounds.

### analyze_cfp_submission

Evaluate conference submissions for content, speaker qualifications and educational value.

### analyze_claims

Evaluate truth claims by analyzing evidence and logical fallacies.

### analyze_comments

Analyze user comments for sentiment, extract praise/criticism, and summarize reception.

### analyze_debate

Analyze debates identifying arguments, agreements, and emotional intensity.

### analyze_interviewer_techniques

Study interviewer questions/methods to identify effective interview techniques.

### analyze_military_strategy

Examine battles analyzing strategic decisions to extract military lessons.

### analyze_mistakes

Analyze past errors to prevent similar mistakes in predictions/decisions.

### analyze_paper

Analyze scientific papers to identify findings and assess conclusion.

### analyze_paper_simple

Analyze research papers to determine primary findings and assess scientific rigor.

### analyze_patent

Analyze patents to evaluate novelty and technical advantages.

### analyze_personality

Psychological analysis by examining language to reveal personality traits.

### analyze_presentation

Evaluate presentations scoring novelty, value for feedback.

### analyze_product_feedback

Process user feedback to identify themes and prioritize insights.

### analyze_proposition

Examine ballot propositions to assess purpose and potential impact.

### analyze_prose

Evaluate writing quality by rating novelty, clarity, and style.

### analyze_prose_json

Evaluate writing and provide JSON output rating novelty, clarity, effectiveness.

### analyze_prose_pinker

Analyze writing style using Pinker's principles to improve clarity and effectiveness.

### analyze_sales_call

Evaluate sales calls analyzing pitch, fundamentals, and customer interaction.

### analyze_spiritual_text

Compare religious texts with KJV, identifying claims and doctrinal variations.

### analyze_tech_impact

Evaluate tech projects' societal impact across dimensions.

### analyze_terraform_plan

Analyze Terraform plans for infrastructure changes, security risks, and cost implications.

### apply_ul_tags

Apply standardized content tags to categorize topics like AI, cybersecurity, politics, and culture.

### check_agreement

Review contract to identify stipulations, issues, and changes for negotiation.

### compare_and_contrast

Create comparisons table, highlighting key differences and similarities.

### concall_summary

Analyze earnings call transcripts to extract management insights, financial metrics, and investment implications.

### create_ai_jobs_analysis

Identify automation risks and career resilience strategies.

### create_better_frame

Develop positive mental frameworks for challenging situations.

### create_story_about_people_interaction

Analyze two personas, compare their dynamics, and craft a realistic, character-driven story from those insights.

### create_idea_compass

Organize thoughts analyzing definitions, evidence, relationships, implications.

### create_recursive_outline

Break down tasks into hierarchical, actionable components via decomposition.

### create_tags

Generate single-word tags for content categorization and mind mapping.

### extract_core_message

Distill the fundamental message into a single, impactful sentence.

### extract_extraordinary_claims

Identify/extract claims contradicting scientific consensus.

### extract_main_idea

Identify key idea, providing core concept and recommendation.

### extract_mcp_servers

Analyzes content to identify and extract detailed information about Model Context Protocol (MCP) servers.

### extract_most_redeeming_thing

Identify the most positive aspect from content.

### extract_predictions

Identify/analyze predictions, claims, confidence, and verification.

### extract_primary_problem

Identify/analyze the core problem / root causes.

### extract_primary_solution

Identify/analyze the main solution proposed in content.

### extract_song_meaning

Analyze song lyrics to uncover deeper meanings and themes.

### find_hidden_message

Analyze content to uncover concealed meanings and implications.

### find_logical_fallacies

Identify/analyze logical fallacies to evaluate argument validity.

### generate_code_rules

Extracts a list of best practices rules for AI coding assisted tools.

### get_wow_per_minute

Calculate frequency of impressive moments to measure engagement.

### identify_dsrp_distinctions

Analyze content using DSRP to identify key distinctions.

### identify_dsrp_perspectives

Analyze content using DSRP to identify different viewpoints.

### identify_dsrp_relationships

Analyze content using DSRP to identify connections.

### identify_dsrp_systems

Analyze content using DSRP to identify systems and structures.

### identify_job_stories

Extract/analyze user job stories to understand motivations.

### label_and_rate

Categorize/evaluate content by assigning labels and ratings.

### model_as_sherlock_freud

Builds psychological models using detective reasoning and psychoanalytic insight.

### predict_person_actions

Predicts behavioral responses based on psychological profiles and challenges

### prepare_7s_strategy

Apply McKinsey 7S framework to analyze organizational alignment.

### provide_guidance

Offer expert advice tailored to situations, providing steps.

### rate_content

Evaluate content quality across dimensions, providing scoring.

### rate_value

Assess practical value of content by evaluating utility.

### recommend_artists

Suggest artists based on user preferences and style.

### recommend_talkpanel_topics

Generate discussion topics for panel talks based on interests.

### summarize_board_meeting

Convert board meeting transcripts into formal meeting notes for corporate records.

### summarize_prompt

Summarize AI prompts to identify instructions and outputs.

### t_analyze_challenge_handling

Evaluate challenge handling by analyzing response strategies.

### t_check_dunning_kruger

Analyze cognitive biases to identify overconfidence and underestimation of abilities using Dunning-Kruger principles.

### t_check_metrics

Analyze metrics, tracking progress and identifying trends.

### t_describe_life_outlook

Analyze personal philosophies to understand core beliefs.

### t_find_blindspots

Identify blind spots in thinking to improve awareness.

### t_find_negative_thinking

Identify negative thinking patterns to recognize distortions.

### t_red_team_thinking

Apply adversarial thinking to identify weaknesses.

### t_year_in_review

Generate annual reviews by analyzing achievements and learnings.

## EXTRACTION PATTERNS

### create_aphorisms

Compile relevant, attributed aphorisms from historical figures on topics.

### create_upgrade_pack

Extract world model updates/algorithms to improve decision-making.

### create_video_chapters

Organize video content into timestamped chapters highlighting key topics.

### extract_algorithm_update_recommendations

Extract recommendations for improving algorithms, focusing on steps.

### extract_alpha

Extracts the most novel and surprising ideas ("alpha") from content, inspired by information theory.

### extract_article_wisdom

Extract wisdom from articles, organizing into actionable takeaways.

### extract_book_ideas

Extract novel ideas from books to inspire new projects.

### extract_book_recommendations

Extract/prioritize practical advice from books.

### extract_characters

Identify all characters (human and non-human), resolve their aliases and pronouns into canonical names, and produce detailed descriptions of each character's role, motivations, and interactions ranked by narrative importance.

### extract_controversial_ideas

Analyze contentious viewpoints while maintaining objective analysis.

### extract_domains

Extract key content and source.

### extract_ideas

Extract/organize concepts and applications into idea collections.

### extract_insights

Extract insights about life, tech, presenting as bullet points.

### extract_insights_dm

Extract insights from DMs, focusing on learnings and takeaways.

### extract_instructions

Extract procedures into clear instructions for implementation.

### extract_latest_video

Extract info from the latest video, including title and content.

### extract_main_activities

Extract and list main events from transcripts.

### extract_patterns

Extract patterns and themes to create reusable templates.

### extract_product_features

Extract/categorize product features into a structured list.

### extract_questions

Extract/categorize questions to create Q&A resources.

### extract_recommendations

Extract recommendations, organizing into actionable guidance.

### extract_references

Extract/format citations into a structured reference list.

### extract_skills

Extract/classify hard/soft skills from job descriptions into skill inventory.

### extract_sponsors

Extract/organize sponsorship info, including names and messages.

### extract_videoid

Extract/parse video IDs and URLs to create video lists.

### extract_wisdom

Extract insightful ideas and recommendations focusing on life wisdom.

### extract_wisdom_dm

Extract learnings from DMs, focusing on personal growth.

### extract_wisdom_nometa

Extract pure wisdom from content without metadata.

### extract_wisdom_short

Extract condensed  insightful ideas and recommendations focusing on life wisdom.

### t_extract_intro_sentences

Extract intro sentences to identify engagement strategies.

### t_extract_panel_topics

Extract panel topics to create engaging discussions.

## SUMMARIZATION PATTERNS

### capture_thinkers_work

Extract key concepts, background, and ideas from notable thinkers' work.

### create_5_sentence_summary

Generate concise summaries of content in five levels, five words to one.

### create_micro_summary

Generate concise summaries with one-sentence overview and key points.

### create_summary

Generate concise summaries by extracting key points and main ideas.

### summarize

Generate summaries capturing key points and details.

### summarize_debate

Summarize debates highlighting arguments and agreements.

### summarize_lecture

Summarize lectures capturing key concepts and takeaways.

### summarize_legislation

Summarize legislation highlighting key provisions and implications.

### summarize_meeting

Summarize meetings capturing discussions and decisions.

### summarize_micro

Generate extremely concise summaries of content.

### summarize_newsletter

Summarize newsletters highlighting updates and trends.

### summarize_paper

Summarize papers highlighting objectives and findings.

### summarize_pull-requests

Summarize pull requests highlighting code changes.

### summarize_rpg_session

Summarize RPG sessions capturing story events and decisions.

### youtube_summary

Summarize YouTube videos with key points and timestamps.

## WRITING PATTERNS

### clean_text

Format/clean text by fixing breaks, punctuation, preserving content/meaning.

### create_academic_paper

Transform content into academic papers using LaTeX layout.

### create_diy

Create step-by-step DIY tutorials with clear instructions and materials.

### create_formal_email

Compose professional emails with proper tone and structure.

### create_keynote

Design TED-style presentations with narrative, slides and notes.

### create_newsletter_entry

Write concise newsletter content focusing on key insights.

### create_show_intro

Craft compelling podcast/show intros to engage audience.

### create_story_about_people_interaction

Analyze two personas, compare their dynamics, and craft a realistic, character-driven story from those insights.

### create_story_explanation

Transform complex concepts into clear, engaging narratives.

### enrich_blog_post

Enhance blog posts by improving structure and visuals for static sites.

### explain_docs

Transform technical docs into clearer explanations with examples.

### explain_terms

Create glossaries of advanced terms with definitions and analogies.

### fix_typos

Proofreads and corrects typos, spelling, grammar, and punctuation errors.

### humanize

Transform technical content into approachable language.

### improve_academic_writing

Enhance academic writing by improving clarity and structure.

### improve_writing

Enhance writing by improving clarity, flow, and style.

### md_callout

Generate markdown callout blocks to highlight info.

### t_create_opening_sentences

Generate compelling opening sentences for content.

### t_give_encouragement

Generate personalized messages of encouragement.

### transcribe_minutes

Convert meeting recordings into structured minutes.

### tweet

Transform content into concise tweets.

### write_essay

Write essays on given topics in the distinctive style of specified authors.

### write_essay_pg

Create essays with thesis statements and arguments in the style of Paul Graham.

### write_latex

Generate LaTeX documents with proper formatting.

### write_micro_essay

Create concise essays presenting a single key idea.

## DEVELOPMENT PATTERNS

### agility_story

Generate agile user stories and acceptance criteria following agile formats.

### analyze_logs

Examine server logs to identify patterns and potential system issues.

### answer_interview_question

Generate appropriate responses to technical interview questions.

### ask_uncle_duke

Expert software dev. guidance focusing on Java, Spring, frontend, and best practices.

### coding_master

Explain coding concepts/languages for beginners

### create_coding_feature

Generate secure and composable code features using latest technology and best practices.

### create_coding_project

Design coding projects with clear architecture, steps, and best practices.

### create_design_document

Create software architecture docs using C4 model.

### create_git_diff_commit

Generate clear git commit messages and commands for code changes.

### create_loe_document

Create detailed Level of Effort (LOE) estimation documents.

### create_prd

Create Product Requirements Documents (PRDs) from input specs.

### create_user_story

Write clear user stories with descriptions and acceptance criteria.

### explain_code

Analyze/explain code, security tool outputs, and configs.

### explain_project

Create project overviews with instructions and usage examples.

### extract_poc

Extract/document proof-of-concept demos from technical content.

### official_pattern_template

Define pattern templates with sections for consistent creation.

### recommend_pipeline_upgrades

Suggest CI/CD pipeline improvements for efficiency and security.

### refine_design_document

Enhance design docs by improving clarity and accuracy.

### review_code

Performs a comprehensive code review, providing detailed feedback on correctness, security, and performance.

### review_design

Evaluate software designs for scalability and security.

### summarize_git_changes

Summarize git changes highlighting key modifications.

### summarize_git_diff

Summarize git diff output highlighting functional changes.

### write_pull-request

Create pull request descriptions with summaries of changes.

## SECURITY PATTERNS

### analyze_email_headers

Analyze email authentication headers to assess security and provide recommendations.

### analyze_incident

Extract info from breach articles, including attack details and impact.

### analyze_malware

Analyze malware behavior, extract IOCs, MITRE ATT&CK, provide recommendations.

### analyze_risk

Assess vendor security compliance to determine risk levels.

### analyze_threat_report

Extract/analyze insights, trends, and recommendations from threat reports.

### analyze_threat_report_cmds

Interpret commands from threat reports, providing implementation guidance.

### analyze_threat_report_trends

Extract/analyze trends from threat reports to identify emerging patterns.

### ask_secure_by_design_questions

Generate security-focused questions to guide secure system design.

### create_command

Generate precise CLI commands for penetration testing tools based on docs.

### create_cyber_summary

Summarize incidents, vulnerabilities into concise intelligence briefings.

### create_network_threat_landscape

Analyze network ports/services to create threat reports with recommendations.

### create_report_finding

Document security findings with descriptions, recommendations, and evidence.

### create_security_update

Compile security newsletters covering threats, advisories, developments with links.

### create_sigma_rules

Extract TTPs and translate them into YAML Sigma detection rules.

### create_stride_threat_model

Generate threat models using STRIDE to prioritize security threats.

### create_threat_scenarios

Develop realistic security threat scenarios based on risk analysis.

### create_ttrc_graph

Generate time-series for visualizing vulnerability remediation metrics.

### create_ttrc_narrative

Create narratives for security program improvements in remediation efficiency.

### extract_ctf_writeup

Extract techniques from CTF writeups to create learning resources.

### improve_report_finding

Enhance security report by improving clarity and accuracy.

### t_threat_model_plans

Analyze plans through a security lens to identify threats.

### write_hackerone_report

Create vulnerability reports following HackerOne's format.

### write_nuclei_template_rule

Generate Nuclei scanning templates with detection logic.

### write_semgrep_rule

Create Semgrep rules for static code analysis.

## BUSINESS PATTERNS

### create_hormozi_offer

Create compelling business offers using Alex Hormozi's methodology.

### extract_business_ideas

Identify business opportunities and insights

### t_create_h3_career

Generate career plans using the Head, Heart, Hands framework.

## LEARNING PATTERNS

### create_flash_cards

Generate flashcards for key concepts and definitions.

### create_quiz

Generate review questions adapting difficulty to student levels.

### create_reading_plan

Design three-phase reading plans to build knowledge of topics.

### dialog_with_socrates

Engage in Socratic dialogue to explore ideas via questioning.

### explain_math

Explain math concepts for students using step-by-step instructions.

### to_flashcards

Convert content into flashcard format for learning.

## VISUALIZATION PATTERNS

### create_conceptmap

Transform unstructured text or markdown content into interactive HTML concept maps using Vis.js by extracting key concepts and their logical relationships.

### create_excalidraw_visualization

Create visualizations using Excalidraw.

### create_graph_from_input

Transform security metrics to CSV for visualizing progress over time.

### create_investigation_visualization

Create Graphviz vis. of investigation data showing relationships and findings.

### create_logo

Generate minimalist logo prompts capturing brand essence via vector graphics.

### create_markmap_visualization

Transform complex ideas into mind maps using Markmap syntax.

### create_mermaid_visualization

Transform concepts into visual diagrams using Mermaid syntax.

### create_mermaid_visualization_for_github

Create Mermaid diagrams to visualize workflows in documentation.

### create_visualization

Transform concepts to ASCII art with explanations of relationships.

### t_visualize_mission_goals_projects

Visualize missions and goals to clarify relationships.

## CONVERSION PATTERNS

### convert_to_markdown

Convert content to markdown, preserving original content and structure.

### export_data_as_csv

Extract data and convert to CSV, preserving data integrity.

### sanitize_broken_html_to_markdown

Clean/convert malformed HTML to markdown.

### translate

Convert content between languages while preserving meaning.

## STRATEGY PATTERNS

### t_find_neglected_goals

Identify neglected goals to surface opportunities.

## PERSONAL DEVELOPMENT PATTERNS

### create_story_about_person

Infer everyday challenges and realistic coping strategies from a psychological profile and craft an empathetic 500–700-word story consistent with the character.

### extract_recipe

Extract/format recipes into instructions with ingredients and steps.

### find_female_life_partner

Clarify and summarize partner criteria in direct language.

### heal_person

Analyze a psychological profile, pinpoint issues and strengths, and deliver compassionate, structured strategies for spiritual, mental, and life improvement.

## CREATIVITY PATTERNS

### create_mnemonic_phrases

Create memorable mnemonic sentences using given words in exact order for memory aids.

## GAMING PATTERNS

### create_npc

Generate detailed D&D 5E NPC characters with backgrounds and game stats.

### create_rpg_summary

Summarize RPG sessions capturing events, combat, and narrative.

## OTHER PATTERNS

### extract_jokes

Extract/categorize jokes, puns, and witty remarks.

## WELLNESS PATTERNS

### recommend_yoga_practice

Provides personalized yoga sequences, meditation guidance, and holistic lifestyle advice based on individual profiles.
