## PDF TO MARKDOWN CONVERSION IMPLEMENTATION

- PDF to Markdown conversion functionality for the web interface
- Automatic detection and processing of PDF files in chat
- Conversion to markdown format for LLM processing
- Installation instructions from the pdf-to-markdown repository

The PDF conversion module has been integrated in the svelte web browser interface. Once installed, it will automatically detect pdf files in the chat interface and convert them to markdown automatically for llm processing.


## HOW TO INSTALL
If you need to update the web component follow the instructions in "Web Interface MOD Readme Files/WEB V2 Install Guide.md".  

Assuming your install is up to date and web svelte config complete, you can simply follow these steps to add Pdf-to-mardown. 

# FROM FABRIC ROOT DIRECTORY
  cd .. web

# Install in this sequence: 
# Step 1
npm install -D patch-package
# Step 2
npm install -D pdfjs-dist@2.5.207
# Step 3
npm install -D github:jzillmann/pdf-to-markdown#modularize


## 🎥 Demo Video (see 4min)
https://youtu.be/bhwtWXoMASA

# Integration with Svelte

The integration approach focused on using the library's high-level API while maintaining SSR compatibility:

- Create PdfConversionService for PDF processing
- Handle file uploads in ChatInput component
- Convert PDF content to markdown text
- Integrate with existing chat processing flow



### How it Works

The PDF to Markdown conversion is implemented as a separate module located in the `pdf-to-markdown` directory. It leverages the `pdf-parse` library (likely via `PdfParser.ts`) to parse PDF documents and extract text content. The core logic resides in `PdfPipeline.ts`, which orchestrates the PDF parsing and conversion process. `Pdf-to-Markdown` is a folk from `pdf.js` - Mozilla's PDF parsing & rendering platform which is used as a raw parser

Here's a simplified breakdown of the process:

1.  **PDF Parsing:** The `PdfParser.ts` uses `pdf-parse` to read the PDF file and extract text content from each page.
2.  **Content Extraction:** The extracted text content is processed to identify text elements, formatting, and structure.
3.  **Markdown Conversion:** The `PdfPipeline.ts` then converts the extracted and processed text content into Markdown format. This involves mapping PDF elements to Markdown syntax, attempting to preserve formatting like headings, lists, and basic text styles.
4.  **Frontend Integration:** The `PdfConversionService.ts` in the `web/src/lib/services` directory acts as a frontend service that utilizes the `pdf-to-markdown` module. It provides a `convertToMarkdown` function that takes a File object (PDF file) as input, calls the `pdf-to-markdown` module to perform the conversion, and returns the Markdown output as a string.
5.  **Chat Input Integration:** The `ChatInput.svelte` component uses the `PdfConversionService` to convert uploaded PDF files to Markdown before sending the content to the chat service for pattern processing.



### File Changes

The following files were added or modified to implement the PDF to Markdown conversion:

**New files:**

*   `pdf-to-markdown/`: (New directory for the PDF to Markdown module)
    *   `pdf-to-markdown/package.json`:  Defines dependencies and build scripts for the PDF to Markdown module.
    *   `pdf-to-markdown/tsconfig.json`: TypeScript configuration for the PDF to Markdown module.
    *   `pdf-to-markdown/src/`: Source code directory for the PDF to Markdown module.
        *   `pdf-to-markdown/src/index.ts`: Entry point of the PDF to Markdown module.
        *   `pdf-to-markdown/src/PdfPipeline.ts`: Core logic for PDF to Markdown conversion pipeline.
        *   `pdf-to-markdown/src/PdfParser.ts`:  PDF parsing logic using `pdf-parse`.

*   `web/src/lib/services/PdfConversionService.ts`: (New file)
    *   Frontend service to use the `pdf-to-markdown` module and expose `convertToMarkdown` function.

**Modified files:**

*   `web/src/lib/components/chat/ChatInput.svelte`:
    *   Modified to import and use the `PdfConversionService` in the `readFileContent` function to handle PDF files.
    *   Modified `readFileContent` to call `pdfService.convertToMarkdown` for PDF files.

These file changes introduce the new PDF to Markdown conversion functionality and integrate it into the chat input component of the web interface.

