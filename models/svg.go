package models

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// SVGResponse represents the response model for SVGHandler
type SVGResponse struct {
	Name1 string `json:"name1"`
	Name2 string `json:"name2"`
}

// KeyValue represents a key-value pair
type KeyValue struct {
	Key   string
	Value int
}

// GenerateSVG generates SVG content based on the names and key-value pairs
func (s SVGResponse) GenerateSVG() string {
    var svgContent strings.Builder
    var iteration_animation int=0

    svgContent.WriteString(`
	<svg id="gh-dark-mode-only" width="380" height="230" xmlns="http://www.w3.org/2000/svg">
	<style>
	svg {
	  font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Helvetica, Arial, sans-serif, Apple Color Emoji, Segoe UI Emoji;
	  font-size: 14px;
	  line-height: 21px;
	}
	
	#background {
	  width: calc(100% - 10px);
	  height: calc(100% - 10px);
	  fill: white;
	  stroke: rgb(225, 228, 232);
	  stroke-width: 1px;
	  rx: 6px;
	  ry: 6px;
	}
	
	#gh-dark-mode-only:target #background {
	  fill: #0d1117;
	  stroke-width: 0.5px;
	}
	
	foreignObject {
	  width: calc(100% - 10px - 32px);
	  height: calc(100% - 10px - 24px);
	}
	
	h2 {
	  margin-top: 0;
	  margin-bottom: 0.75em;
	  line-height: 24px;
	  font-size: 16px;
	  font-weight: 600;
	  color: rgb(36, 41, 46);
	  fill: rgb(36, 41, 46);
	}
	
	#gh-dark-mode-only:target h2 {
	  color: #c9d1d9;
	  fill: #c9d1d9;
	}
	
	ul {
	  list-style: none;
	  padding-left: 0;
	  margin-top: 0;
	  margin-bottom: 0;
	}
	
	li {
	  display: inline-flex;
	  font-size: 12px;
	  margin-right: 2ch;
	  align-items: center;
	  flex-wrap: nowrap;
	  transform: translateX(-500%);
	  animation: slideIn 2s ease-in-out forwards;
	}
	
	@keyframes slideIn {
	  to {
		transform: translateX(0);
	  }
	}
	
	div.ellipsis {
	  height: 100%;
	  overflow: hidden;
	  text-overflow: ellipsis;
	}
	
	.octicon {
	  fill: rgb(88, 96, 105);
	  margin-right: 0.5ch;
	  vertical-align: top;
	}
	
	#gh-dark-mode-only:target .octicon {
	  color: #8b949e;
	  fill: #8b949e;
	}
	
	.progress {
	  display: flex;
	  height: 8px;
	  overflow: hidden;
	  background-color: rgb(225, 228, 232);
	  border-radius: 6px;
	  outline: 1px solid transparent;
	  margin-bottom: 1em;
	}
	
	#gh-dark-mode-only:target .progress {
	  background-color: rgba(110, 118, 129, 0.4);
	}
	
	.progress-item {
	  outline: 2px solid rgb(225, 228, 232);
	  border-collapse: collapse;
	}
	
	#gh-dark-mode-only:target .progress-item {
	  outline: 2px solid #393f47;
	}
	
	.lang {
	  font-weight: 600;
	  margin-right: 4px;
	  color: rgb(36, 41, 46);
	}
	
	#gh-dark-mode-only:target .lang {
	  color: #c9d1d9;
	}
	
	.percent {
	  color: rgb(88, 96, 105)
	}
	
	#gh-dark-mode-only:target .percent {
	  color: #8b949e;
	}
	</style>
	<g>
	<rect x="5" y="5" id="background"/>
	<g>
	<foreignObject x="21" y="17" width="318" height="176">
	<div xmlns="http://www.w3.org/1999/xhtml" class="ellipsis">`)

    svgContent.WriteString(fmt.Sprintf(`<h2>%s</h2>

	<div>
	<span class="progress">`, s.Name1))


	
    // Parse the second variable into an array of key-value pairs
    keyValuePairs, errors := parseKeyValuePairs(s.Name2)

    // Calculate the sum of all num2 values
    sum := 0
    for _, kv := range keyValuePairs {
        sum += kv.Value
    }

	colors := []string{
        "#FF6347", // Vivid Sunset Orange
        "#FFC0CB", // Soft Blush Pink
        "#ADD8E6", // Tranquil Light Blue
        "#90EE90", // Serene Spring Green
        "#F0E68C", // Warm Sunshine Yellow
        "#E0FFFF", // Calming Sky Blue
        "#AFEEEE", // Refreshing Mint Green
        "#D3959B", // Delicate Lavender
        "#F4C2C2", // Elegant Dusty Rose
        "#98FB98", // Calming Seafoam Green
        "#FFA500", // Energetic Orange
        "#800080", // Deep Purple
        "#C71585", // Bold Magenta
        "#00CED1", // Soothing Turquoise
        "#FFFF00", // Bright Yellow
        "#DC143C", // Deep Crimson Red
        "#000000", // Classic Black
        "#FFFFFF", // Pure White
        "#A52A2A", // Earthy Brown
        "#7FFFD4", // Gentle Aqua Blue
    }
	
    // Add ratios to the SVG content
    y := 50
    for i, kv := range keyValuePairs {
        ratio := float64(kv.Value) / float64(sum) * 100
        svgContent.WriteString(fmt.Sprintf(`<span style="background-color: %s;width: %.2f%%;" class="progress-item"></span>`,colors[i], ratio)) //, y , kv.Key
        y += 20
    }
	
	svgContent.WriteString(`</span>
	</div>
	
	<ul>`)



    // Display original values
    y += 20

    for i, kv := range keyValuePairs {
	
        svgContent.WriteString(fmt.Sprintf(`
	
 <li style="animation-delay: %d ms;">
		<svg xmlns="http://www.w3.org/2000/svg" class="octicon" style="fill:%s;" viewBox="0 0 16 16" version="1.1" width="16" height="16"><path fill-rule="evenodd" d="M8 4a4 4 0 100 8 4 4 0 000-8z"/></svg>
		<span class="lang">%s</span>
		<span class="percent">%d.0%%</span>
		</li>
`,iteration_animation,colors[i],kv.Key,kv.Value))
        y += 20
        iteration_animation += 150
    }








    svgContent.WriteString(`
	</ul>
	
	</div>
	</foreignObject>
	</g>
	</g>
	</svg>`)


    // Log any errors encountered during parsing
    for _, err := range errors {
        log.Println("Error:", err)
    }

    return svgContent.String()
}




// parseKeyValuePairs parses the string of key-value pairs separated by commas into an array of KeyValue structs
func parseKeyValuePairs(input string) ([]KeyValue, []error) {
    keyValuePairs := strings.Split(input, ",")
    result := make([]KeyValue, 0, len(keyValuePairs))
    errors := make([]error, 0)

    for _, pair := range keyValuePairs {
        parts := strings.Split(pair, "=")

        // Skip empty or malformed pairs
        if len(parts) != 2 {
            errors = append(errors, fmt.Errorf("malformed pair: %s", pair))
            continue
        }

        // Parse the key
        key := strings.TrimSpace(parts[0])

        // Parse the value
        valueStr := strings.TrimSpace(parts[1])
        value, err := strconv.Atoi(valueStr)
        if err != nil {
            // If parsing the value fails, add an error to the list
            errors = append(errors, fmt.Errorf("failed to parse value '%s' as an integer: %v", valueStr, err))
            continue
        }

        // Check if key or value is empty
        if key == "" || valueStr == "" {
            errors = append(errors, fmt.Errorf("empty key or value in pair: %s", pair))
            continue
        }

        // Add the key-value pair to the result
        result = append(result, KeyValue{Key: key, Value: value})
    }

    return result, errors
}
