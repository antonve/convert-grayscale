on run {input, parameters}
	set arguments to " "
	repeat with anImage in input
		set arguments to arguments & " -f " & (quoted form of POSIX path of anImage)
	end repeat
	
	do shell script "/Users/anton/xdev/go/src/github.com/antonve/convert-gray/convert-gray" & arguments
end run