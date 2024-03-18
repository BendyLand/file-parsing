def split_into_blocks(lines)
    blocks = []
    block = ""
    lines.each do |line|
        if line[0] != " "
            blocks.push(block)
            block = ""
            block += line + "\t"
        else
            block += line
        end
    end
    return blocks
end

def format_block(block) 
    result = ""
    lines = block.split("\t")

    language = lines[0]
    details = lines[1].split(" ")
    formatted_lines = []
    detail_str = ""
    details.each do |detail|
        if detail.include?(":")
            formatted_lines.push(detail_str)
            detail_str = ""
            detail_str += detail
        else
            detail_str += " " + detail
        end
    end
    return language + formatted_lines.join("\n") + "\n\n"
end

def extract_language_type(block)
    lines = block.split("\n")
    return [lines[0], lines[1]]
end

file = File.open("../languages-on-github.yml")
lines = file.readlines.map(&:chomp).filter { |line| line[0] != "#" }
blocks = split_into_blocks(lines).filter { |line| line.length > 10 }
formatted_blocks = blocks.map { |block| format_block(block) }
lang_types = formatted_blocks.map { |block| extract_language_type(block) }
programming_langs = lang_types.filter { |lang| lang[1] == "type: programming" }
non_programming_langs = lang_types.filter { |lang| lang[1] != "type: programming" }

puts "Programming languages:"
programming_langs.each do |lang|
    puts lang[0][..-2]
end
puts "\nNon-programming languages:"
non_programming_langs.each do |lang|
    puts lang[0][..-2]
end
