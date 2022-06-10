// Package namegen generates random names.
package namegen

import (
	"math/rand"
	"strconv"
	"time"
)

var (
	left = [...]string{
		"able",
		"accurate",
		"actual",
		"additional",
		"admiring",
		"adoring",
		"affectionate",
		"afraid",
		"aggressive",
		"agitated",
		"alive",
		"amazing",
		"angry",
		"anxious",
		"asleep",
		"attentive",
		"automatic",
		"available",
		"aware",
		"awesome",
		"bad",
		"basic",
		"beautiful",
		"best",
		"big",
		"blissful",
		"bold",
		"boring",
		"brave",
		"bright",
		"busy",
		"calm",
		"capable",
		"careful",
		"central",
		"certain",
		"charming",
		"cheap",
		"civil",
		"clean",
		"clear",
		"clever",
		"close",
		"cold",
		"comfortable",
		"common",
		"compassionate",
		"competent",
		"condescending",
		"confident",
		"conscious",
		"consistent",
		"constant",
		"convenient",
		"cool",
		"correct",
		"cranky",
		"crazy",
		"critical",
		"cultural",
		"curious",
		"current",
		"dangerous",
		"dark",
		"dazzling",
		"dead",
		"decent",
		"deep",
		"democratic",
		"desperate",
		"determined",
		"different",
		"difficult",
		"dirty",
		"distinct",
		"distracted",
		"dreamy",
		"eager",
		"early",
		"eastern",
		"easy",
		"economic",
		"ecstatic",
		"educational",
		"efficient",
		"elastic",
		"elated",
		"elegant",
		"eloquent",
		"emotional",
		"empty",
		"entire",
		"epic",
		"every",
		"exact",
		"exciting",
		"existing",
		"expensive",
		"fair",
		"famous",
		"fast",
		"fat",
		"fervent",
		"festive",
		"final",
		"financial",
		"fine",
		"firm",
		"flamboyant",
		"flat",
		"focused",
		"foreign",
		"formal",
		"former",
		"free",
		"fresh",
		"friendly",
		"frightful",
		"frosty",
		"full",
		"funny",
		"gallant",
		"general",
		"gifted",
		"global",
		"good",
		"goofy",
		"gorgeous",
		"gracious",
		"great",
		"guilty",
		"happy",
		"hard",
		"hardcore",
		"healthy",
		"heavy",
		"helpful",
		"heuristic",
		"high",
		"historical",
		"honest",
		"hopeful",
		"hot",
		"huge",
		"human",
		"hungry",
		"ill",
		"illegal",
		"immediate",
		"important",
		"impossible",
		"impressive",
		"independent",
		"infallible",
		"informal",
		"inner",
		"innocent",
		"inspiring",
		"intelligent",
		"interesting",
		"international",
		"jolly",
		"jovial",
		"keen",
		"kind",
		"known",
		"large",
		"late",
		"laughing",
		"leading",
		"left",
		"legal",
		"light",
		"likely",
		"little",
		"local",
		"lonely",
		"long",
		"loose",
		"loud",
		"loving",
		"low",
		"lucid",
		"lucky",
		"magical",
		"main",
		"major",
		"medical",
		"mental",
		"military",
		"modest",
		"musing",
		"mystifying",
		"national",
		"natural",
		"naughty",
		"necessary",
		"nervous",
		"new",
		"nice",
		"nifty",
		"normal",
		"nostalgic",
		"numerous",
		"objective",
		"obvious",
		"official",
		"old",
		"open",
		"opposite",
		"optimistic",
		"past",
		"peaceful",
		"pedantic",
		"pensive",
		"perfect",
		"personal",
		"physical",
		"pleasant",
		"polite",
		"political",
		"poor",
		"popular",
		"possible",
		"powerful",
		"practical",
		"priceless",
		"private",
		"psychological",
		"public",
		"pure",
		"quiet",
		"quirky",
		"quizzical",
		"rare",
		"ready",
		"real",
		"realistic",
		"reasonable",
		"recent",
		"recursing",
		"relaxed",
		"relevant",
		"religious",
		"remarkable",
		"remote",
		"responsible",
		"reverent",
		"rich",
		"right",
		"romantic",
		"rude",
		"sad",
		"safe",
		"scared",
		"secure",
		"sensible",
		"serene",
		"serious",
		"severe",
		"sexual",
		"sharp",
		"short",
		"significant",
		"silly",
		"similar",
		"simple",
		"single",
		"sleepy",
		"slow",
		"small",
		"smooth",
		"social",
		"sorry",
		"special",
		"stoic",
		"strange",
		"strict",
		"strong",
		"stupefied",
		"substantial",
		"successful",
		"sudden",
		"sufficient",
		"suitable",
		"sure",
		"suspicious",
		"sweet",
		"tall",
		"tasty",
		"technical",
		"tender",
		"terrible",
		"thin",
		"thirsty",
		"tight",
		"tiny",
		"tired",
		"traditional",
		"true",
		"trusting",
		"typical",
		"ugly",
		"unable",
		"unfair",
		"unhappy",
		"united",
		"unlikely",
		"unruffled",
		"unusual",
		"upbeat",
		"used",
		"useful",
		"usual",
		"valuable",
		"various",
		"vibrant",
		"vigilant",
		"vigorous",
		"visible",
		"warm",
		"weak",
		"weird",
		"whole",
		"wide",
		"willing",
		"wise",
		"wizardly",
		"wonderful",
		"wrong",
		"xenodochial",
		"young",
		"youthful",
		"zealous",
		"zen",
	}

	// Docker, starting from 0.7.x, generates names from notable scientists and hackers.
	// Please, for any amazing man that you add to the list, consider adding an equally amazing woman to it, and vice versa.
	right = [...]string{

		// Ameyo Stella Shade Adadevoh - Nigerian physician whose expertise and heroic efforts curbed the spread of Ebola in Nigeria in 2014. https://en.wikipedia.org/wiki/Ameyo_Adadevoh
		"adadevoh",

		// Maria Gaetana Agnesi - Italian mathematician, philosopher, theologian and humanitarian. She was the first woman to write a mathematics handbook and the first woman appointed as a Mathematics Professor at a University. https://en.wikipedia.org/wiki/Maria_Gaetana_Agnesi
		"agnesi",

		// Agnodice - First female midwife or physician in ancient Athens. https://en.wikipedia.org/wiki/Agnodice
		"agnodice",

		// Muhammad ibn Jābir al-Ḥarrānī al-Battānī was a founding father of astronomy. https://en.wikipedia.org/wiki/Mu%E1%B8%A5ammad_ibn_J%C4%81bir_al-%E1%B8%A4arr%C4%81n%C4%AB_al-Batt%C4%81n%C4%AB
		"albattani",

		// Frances E. Allen, became the first female IBM Fellow in 1989. In 2006, she became the first female recipient of the ACM's Turing Award. https://en.wikipedia.org/wiki/Frances_E._Allen
		"allen",

		// June Almeida - Scottish virologist who took the first pictures of the rubella virus - https://en.wikipedia.org/wiki/June_Almeida
		"almeida",

		// Kathleen Antonelli, American computer programmer and one of the six original programmers of the ENIAC - https://en.wikipedia.org/wiki/Kathleen_Antonelli
		"antonelli",

		// Virginia Apgar - an American anesthesiologist, the inventor of the Apgar Score, a way to quickly assess the health of a newborn child. https://en.wikipedia.org/wiki/Virginia_Apgar
		"apgar",

		// Archimedes was a physicist, engineer and mathematician who invented too many things to list them here. https://en.wikipedia.org/wiki/Archimedes
		"archimedes",

		// Maria Ardinghelli - Italian translator, mathematician and physicist - https://en.wikipedia.org/wiki/Maria_Ardinghelli
		"ardinghelli",

		// Neil Armstrong - American astronaut, aeronautical engineer, and the first person to walk on the Moon. - https://en.wikipedia.org/wiki/Neil_Armstrong
		"armstrong",

		// Aryabhata - Ancient Indian mathematician-astronomer during 476-550 CE https://en.wikipedia.org/wiki/Aryabhata
		"aryabhata",

		// Barbara S. Askins - American chemist who invented a method to enhance underexposed photographic negatives (used by NASA and the medical industry). https://en.wikipedia.org/wiki/Barbara_Askins
		"askins",

		// Wanda Austin - Wanda Austin is the President and CEO of The Aerospace Corporation, a leading architect for the US security space programs. https://en.wikipedia.org/wiki/Wanda_Austin
		"austin",

		// Charles Babbage invented the concept of a programmable computer. https://en.wikipedia.org/wiki/Charles_Babbage.
		"babbage",

		// Alice Augusta Ball - African American chemist who developed an injectable herbal extract (ethyl hydnocarpate) that was the most effective treatment for leprosy during the early 20th century. https://en.wikipedia.org/wiki/Alice_Ball
		"ball",

		// Stefan Banach - Polish mathematician, was one of the founders of modern functional analysis. https://en.wikipedia.org/wiki/Stefan_Banach
		"banach",

		// Buckaroo Banzai and his mentor Dr. Hikita perfected the "oscillation overthruster", a device that allows one to pass through solid matter. - https://en.wikipedia.org/wiki/The_Adventures_of_Buckaroo_Banzai_Across_the_8th_Dimension
		"banzai",

		// John Bardeen co-invented the transistor - https://en.wikipedia.org/wiki/John_Bardeen
		"bardeen",

		// Jean Bartik, born Betty Jean Jennings, was one of the original programmers for the ENIAC computer. https://en.wikipedia.org/wiki/Jean_Bartik
		"bartik",

		// Clara Barton - Founder of the American Red Cross, hospital nurse during the American Civil War, and civil rights advocate. https://en.wikipedia.org/wiki/Clara_Barton
		"barton",

		// Laura Bassi, the world's first female professor https://en.wikipedia.org/wiki/Laura_Bassi
		"bassi",

		// Hugh Beaver, British engineer, founder of the Guinness Book of World Records https://en.wikipedia.org/wiki/Hugh_Beaver
		"beaver",

		// Alexander Graham Bell - an eminent Scottish-born scientist, inventor, engineer and innovator who is credited with inventing the first practical telephone - https://en.wikipedia.org/wiki/Alexander_Graham_Bell
		"bell",

		// Karl Friedrich Benz - a German automobile engineer. Inventor of the first practical motorcar. https://en.wikipedia.org/wiki/Karl_Benz
		"benz",

		// Homi J Bhabha - was an Indian nuclear physicist, founding director, and professor of physics at the Tata Institute of Fundamental Research. Colloquially known as "father of Indian nuclear programme"- https://en.wikipedia.org/wiki/Homi_J._Bhabha
		"bhabha",

		// Bhaskara II - Ancient Indian mathematician-astronomer whose work on calculus predates Newton and Leibniz by over half a millennium - https://en.wikipedia.org/wiki/Bh%C4%81skara_II#Calculus
		"bhaskara",

		// Sue Black - British computer scientist and campaigner. She has been instrumental in saving Bletchley Park, the site of World War II codebreaking - https://en.wikipedia.org/wiki/Sue_Black_(computer_scientist)
		"black",

		// Elizabeth Helen Blackburn - Australian-American Nobel laureate; best known for co-discovering telomerase. https://en.wikipedia.org/wiki/Elizabeth_Blackburn
		"blackburn",

		// Elizabeth Blackwell - American doctor and first American woman to receive a medical degree - https://en.wikipedia.org/wiki/Elizabeth_Blackwell
		"blackwell",

		// Susanne Bødker - Researcher and author focued on human-computer interaction https://en.wikipedia.org/wiki/Susanne_B%C3%B8dker
		"bodker",

		// Niels Bohr is the father of quantum theory. https://en.wikipedia.org/wiki/Niels_Bohr.
		"bohr",

		// Kathleen Booth, she's credited with writing the first assembly language. https://en.wikipedia.org/wiki/Kathleen_Booth
		"booth",

		// Anita Borg - Anita Borg was the founding director of the Institute for Women and Technology (IWT). https://en.wikipedia.org/wiki/Anita_Borg
		"borg",

		// Satyendra Nath Bose - He provided the foundation for Bose–Einstein statistics and the theory of the Bose–Einstein condensate. - https://en.wikipedia.org/wiki/Satyendra_Nath_Bose
		"bose",

		// Katherine Louise Bouman is an imaging scientist and Assistant Professor of Computer Science at the California Institute of Technology. She researches computational methods for imaging, and developed an algorithm that made possible the picture first visualization of a black hole using the Event Horizon Telescope. - https://en.wikipedia.org/wiki/Katie_Bouman
		"bouman",

		// Evelyn Boyd Granville - She was one of the first African-American woman to receive a Ph.D. in mathematics; she earned it in 1949 from Yale University. https://en.wikipedia.org/wiki/Evelyn_Boyd_Granville
		"boyd",

		// Brahmagupta - Ancient Indian mathematician during 598-670 CE who gave rules to compute with zero - https://en.wikipedia.org/wiki/Brahmagupta#Zero
		"brahmagupta",

		// Walter Houser Brattain co-invented the transistor - https://en.wikipedia.org/wiki/Walter_Houser_Brattain
		"brattain",

		// Harriet Brooks - Canadian Nuclear Physicist. Famous for her research on nuclear transmutations and radioactivity. Among the first persons to discover radon and to try to determine its atomic mass. https://en.wikipedia.org/wiki/Harriet_Brooks
		"brooks",

		// Emmett Brown invented time travel. https://en.wikipedia.org/wiki/Emmett_Brown (thanks Brian Goff)
		"brown",

		// Linda Brown Buck - American biologist and Nobel laureate best known for her genetic and molecular analyses of the mechanisms of smell. https://en.wikipedia.org/wiki/Linda_B._Buck
		"buck",

		// Dame Susan Jocelyn Bell Burnell - Northern Irish astrophysicist who discovered radio pulsars and was the first to analyse them. https://en.wikipedia.org/wiki/Jocelyn_Bell_Burnell
		"burnell",

		// Annie Jump Cannon - pioneering female astronomer who classified hundreds of thousands of stars and created the system we use to understand stars today. https://en.wikipedia.org/wiki/Annie_Jump_Cannon
		"cannon",

		// Rachel Carson - American marine biologist and conservationist, her book Silent Spring and other writings are credited with advancing the global environmental movement. https://en.wikipedia.org/wiki/Rachel_Carson
		"carson",

		// Dame Mary Lucy Cartwright - British mathematician who was one of the first to study what is now known as chaos theory. Also known for Cartwright's theorem which finds applications in signal processing. https://en.wikipedia.org/wiki/Mary_Cartwright
		"cartwright",

		// George Washington Carver - American agricultural scientist and inventor. He was the most prominent black scientist of the early 20th century. https://en.wikipedia.org/wiki/George_Washington_Carver
		"carver",

		// Vinton Gray Cerf - American Internet pioneer, recognised as one of "the fathers of the Internet". With Robert Elliot Kahn, he designed TCP and IP, the primary data communication protocols of the Internet and other computer networks. https://en.wikipedia.org/wiki/Vint_Cerf
		"cerf",

		// Subrahmanyan Chandrasekhar - Astrophysicist known for his mathematical theory on different stages and evolution in structures of the stars. He has won nobel prize for physics - https://en.wikipedia.org/wiki/Subrahmanyan_Chandrasekhar
		"chandrasekhar",

		// Sergey Alexeyevich Chaplygin (Russian: Серге́й Алексе́евич Чаплы́гин; April 5, 1869 – October 8, 1942) was a Russian and Soviet physicist, mathematician, and mechanical engineer. He is known for mathematical formulas such as Chaplygin's equation and for a hypothetical substance in cosmology called Chaplygin gas, named after him. https://en.wikipedia.org/wiki/Sergey_Chaplygin
		"chaplygin",

		// Émilie du Châtelet - French natural philosopher, mathematician, physicist, and author during the early 1730s, known for her translation of and commentary on Isaac Newton's book Principia containing basic laws of physics. https://en.wikipedia.org/wiki/%C3%89milie_du_Ch%C3%A2telet
		"chatelet",

		// Asima Chatterjee was an Indian organic chemist noted for her research on vinca alkaloids, development of drugs for treatment of epilepsy and malaria - https://en.wikipedia.org/wiki/Asima_Chatterjee
		"chatterjee",

		// David Lee Chaum - American computer scientist and cryptographer. Known for his seminal contributions in the field of anonymous communication. https://en.wikipedia.org/wiki/David_Chaum
		"chaum",

		// Pafnuty Chebyshev - Russian mathematician. He is known fo his works on probability, statistics, mechanics, analytical geometry and number theory https://en.wikipedia.org/wiki/Pafnuty_Chebyshev
		"chebyshev",

		// Joan Clarke - Bletchley Park code breaker during the Second World War who pioneered techniques that remained top secret for decades. Also an accomplished numismatist https://en.wikipedia.org/wiki/Joan_Clarke
		"clarke",

		// Bram Cohen - American computer programmer and author of the BitTorrent peer-to-peer protocol. https://en.wikipedia.org/wiki/Bram_Cohen
		"cohen",

		// Jane Colden - American botanist widely considered the first female American botanist - https://en.wikipedia.org/wiki/Jane_Colden
		"colden",

		// Fernando Corbató - American computer scientist instrumental in developing the first time-sharing systems (which directly inspired Unix). https://en.wikipedia.org/wiki/Fernando_J._Corbat%C3%B3
		"corbato",

		// Nicolaus Copernicus - Renaissance-era polymath, the author of theory of the universe with the Sun rather than Earth at the center. https://en.wikipedia.org/wiki/Nicolaus_Copernicus
		"copernicus",

		// Gerty Theresa Cori - American biochemist who became the third woman—and first American woman—to win a Nobel Prize in science, and the first woman to be awarded the Nobel Prize in Physiology or Medicine. Cori was born in Prague. https://en.wikipedia.org/wiki/Gerty_Cori
		"cori",

		// Thomas H. Cormen - American computer scientist and professor. Co-author of the book Introduction To Algorithms. https://en.wikipedia.org/wiki/Thomas_H._Cormen
		"cormen",

		// Seymour Roger Cray was an American electrical engineer and supercomputer architect who designed a series of computers that were the fastest in the world for decades. https://en.wikipedia.org/wiki/Seymour_Cray
		"cray",

		// This entry reflects a husband and wife team who worked together:
		// Joan Curran was a Welsh scientist who developed radar and invented chaff, a radar countermeasure. https://en.wikipedia.org/wiki/Joan_Curran
		// Samuel Curran was an Irish physicist who worked alongside his wife during WWII and invented the proximity fuse. https://en.wikipedia.org/wiki/Samuel_Curran
		"curran",

		// Marie Curie discovered radioactivity. https://en.wikipedia.org/wiki/Marie_Curie.
		"curie",

		// Joan Daemen - Belgian cryptographer best known as the co-designer of Rijndael, which went on to be selected as the Advanced Encryption Standard (AES) and Keccak, which went on to be selected as the Secure Hash Algorithm 3 (SHA-3). https://en.wikipedia.org/wiki/Joan_Daemen
		"daemen",

		// Charles Darwin established the principles of natural evolution. https://en.wikipedia.org/wiki/Charles_Darwin.
		"darwin",

		// Leonardo Da Vinci invented too many things to list here. https://en.wikipedia.org/wiki/Leonardo_da_Vinci.
		"davinci",

		// Democritus - a pre-socratic philosopher who was one of the founders of the Atomist school of philosophy. https://en.wikipedia.org/wiki/Democritus
		"democritus",

		// A. K. (Alexander Keewatin) Dewdney, Canadian mathematician, computer scientist, author and filmmaker. Contributor to Scientific American's "Computer Recreations" from 1984 to 1991. Author of Core War (program), The Planiverse, The Armchair Universe, The Magic Machine, The New Turing Omnibus, and more. https://en.wikipedia.org/wiki/Alexander_Dewdney
		"dewdney",

		// Satish Dhawan - Indian mathematician and aerospace engineer, known for leading the successful and indigenous development of the Indian space programme. https://en.wikipedia.org/wiki/Satish_Dhawan
		"dhawan",

		// Bailey Whitfield Diffie - American cryptographer and one of the pioneers of public-key cryptography. https://en.wikipedia.org/wiki/Whitfield_Diffie
		"diffie",

		// Edsger Wybe Dijkstra was a Dutch computer scientist and mathematical scientist. https://en.wikipedia.org/wiki/Edsger_W._Dijkstra.
		"dijkstra",

		// Paul Adrien Maurice Dirac - English theoretical physicist who made fundamental contributions to the early development of both quantum mechanics and quantum electrodynamics. https://en.wikipedia.org/wiki/Paul_Dirac
		"dirac",

		// Agnes Meyer Driscoll - American cryptanalyst during World Wars I and II who successfully cryptanalysed a number of Japanese ciphers. She was also the co-developer of one of the cipher machines of the US Navy, the CM. https://en.wikipedia.org/wiki/Agnes_Meyer_Driscoll
		"driscoll",

		// Donna Dubinsky - played an integral role in the development of personal digital assistants (PDAs) serving as CEO of Palm, Inc. and co-founding Handspring. https://en.wikipedia.org/wiki/Donna_Dubinsky
		"dubinsky",

		// Sylvia Earle - American oceanographer and marine biologist known for her research on marine algae and her books and documentaries designed to raise awareness of the threats that overfishing and pollution pose to the world’s oceans. https://en.wikipedia.org/wiki/Sylvia_Earle
		"earle",

		// Annie Easley - She was a leading member of the team which developed software for the Centaur rocket stage and one of the first African-Americans in her field. https://en.wikipedia.org/wiki/Annie_Easley
		"easley",

		// Thomas Alva Edison, prolific inventor https://en.wikipedia.org/wiki/Thomas_Edison
		"edison",

		// Albert Einstein invented the general theory of relativity. https://en.wikipedia.org/wiki/Albert_Einstein
		"einstein",

		// Alexandra Asanovna Elbakyan (Russian: Алекса́ндра Аса́новна Элбакя́н) is a Kazakhstani graduate student, computer programmer, internet pirate in hiding, and the creator of the site Sci-Hub. Nature has listed her in 2016 in the top ten people that mattered in science, and Ars Technica has compared her to Aaron Swartz. - https://en.wikipedia.org/wiki/Alexandra_Elbakyan
		"elbakyan",

		// Taher A. ElGamal - Egyptian cryptographer best known for the ElGamal discrete log cryptosystem and the ElGamal digital signature scheme. https://en.wikipedia.org/wiki/Taher_Elgamal
		"elgamal",

		// Gertrude Elion - American biochemist, pharmacologist and the 1988 recipient of the Nobel Prize in Medicine - https://en.wikipedia.org/wiki/Gertrude_Elion
		"elion",

		// James Henry Ellis - British engineer and cryptographer employed by the GCHQ. Best known for conceiving for the first time, the idea of public-key cryptography. https://en.wikipedia.org/wiki/James_H._Ellis
		"ellis",

		// Douglas Engelbart gave the mother of all demos: https://en.wikipedia.org/wiki/Douglas_Engelbart
		"engelbart",

		// Euclid invented geometry. https://en.wikipedia.org/wiki/Euclid
		"euclid",

		// Leonhard Euler invented large parts of modern mathematics. https://de.wikipedia.org/wiki/Leonhard_Euler
		"euler",

		// Michael Faraday - British scientist who contributed to the study of electromagnetism and electrochemistry. https://en.wikipedia.org/wiki/Michael_Faraday
		"faraday",

		// Marga Faulstich - German glass chemist known for her research which forms the basis of techniques used in the manufacture of lightweight anti-reflective glasses; inventor of the SF 64 corrective lens (which was thinner, lighter and more aesthetically pleasing). https://en.wikipedia.org/wiki/Marga_Faulstich
		"faulstich",

		// Genevieve Grotjan Feinstein - American mathematician and cryptanalyst known for deciphering the Japanese World War II cipher machine, Purple and her later work on the Cold War-era Venona project. https://en.wikipedia.org/wiki/Genevieve_Grotjan_Feinstein
		"feinstein",

		// Horst Feistel - German-born American cryptographer who was one of the earliest non-government researchers to study the design and theory of block ciphers. Co-developer of DES and Lucifer. Feistel networks, a symmetric structure used in the construction of block ciphers are named after him. https://en.wikipedia.org/wiki/Horst_Feistel
		"feistel",

		// Pierre de Fermat pioneered several aspects of modern mathematics. https://en.wikipedia.org/wiki/Pierre_de_Fermat
		"fermat",

		// Enrico Fermi invented the first nuclear reactor. https://en.wikipedia.org/wiki/Enrico_Fermi.
		"fermi",

		// Richard Feynman was a key contributor to quantum mechanics and particle physics. https://en.wikipedia.org/wiki/Richard_Feynman
		"feynman",

		// Benjamin Franklin is famous for his experiments in electricity and the invention of the lightning rod.
		"franklin",

		// Elizebeth Smith Friedman was dubbed as "America's first female cryptanalyst". During World War II she solved numerous cipher systems used by the Germans. https://en.wikipedia.org/wiki/Elizebeth_Smith_Friedman
		"friedman",

		// Yuri Alekseyevich Gagarin - Soviet pilot and cosmonaut, best known as the first human to journey into outer space. https://en.wikipedia.org/wiki/Yuri_Gagarin
		"gagarin",

		// Galileo was a founding father of modern astronomy, and faced politics and obscurantism to establish scientific truth.  https://en.wikipedia.org/wiki/Galileo_Galilei
		"galileo",

		// Évariste Galois - French mathematician whose work laid the foundations of Galois theory and group theory, two major branches of abstract algebra, and the subfield of Galois connections, all while still in his late teens. https://en.wikipedia.org/wiki/%C3%89variste_Galois
		"galois",

		// Kadambini Ganguly - Indian physician, known for being the first South Asian female physician, trained in western medicine, to graduate in South Asia. https://en.wikipedia.org/wiki/Kadambini_Ganguly
		"ganguly",

		// William Henry "Bill" Gates III is an American business magnate, philanthropist, investor, computer programmer, and inventor. https://en.wikipedia.org/wiki/Bill_Gates
		"gates",

		// Johann Carl Friedrich Gauss - German mathematician who made significant contributions to many fields, including number theory, algebra, statistics, analysis, differential geometry, geodesy, geophysics, mechanics, electrostatics, magnetic fields, astronomy, matrix theory, and optics. https://en.wikipedia.org/wiki/Carl_Friedrich_Gauss
		"gauss",

		// Maiju Gebhard - developed and popularized the dish drying cabinet which have become a standard accessory in virtually every Finnish home saving approximately 50% of the time related to doing dishes. https://en.wikipedia.org/wiki/Dish_drying_cabinet
		"gebhard",

		// Marie-Sophie Germain - French mathematician, physicist and philosopher. Known for her work on elasticity theory, number theory and philosophy. https://en.wikipedia.org/wiki/Sophie_Germain
		"germain",

		// Maria Goeppert Mayer - German-American theoretical physicist and Nobel Laureate in Physics. Known for her discovery of the nuclear shell model of the atomic nucleus. https://en.wikipedia.org/wiki/Maria_Goeppert_Mayer
		"goeppert",

		// Adele Goldberg, was one of the designers and developers of the Smalltalk language. https://en.wikipedia.org/wiki/Adele_Goldberg_(computer_scientist)
		"goldberg",

		// Adele Goldstine, born Adele Katz, wrote the complete technical description for the first electronic digital computer, ENIAC. https://en.wikipedia.org/wiki/Adele_Goldstine
		"goldstine",

		// Shafi Goldwasser is a computer scientist known for creating theoretical foundations of modern cryptography. Winner of 2012 ACM Turing Award. https://en.wikipedia.org/wiki/Shafi_Goldwasser
		"goldwasser",

		// James Golick, all around gangster.
		"golick",

		// Jane Goodall - British primatologist, ethologist, and anthropologist who is considered to be the world's foremost expert on chimpanzees - https://en.wikipedia.org/wiki/Jane_Goodall
		"goodall",

		//James Arthur Gosling, often referred to as "Dr. Java", OC (born May 19, 1955) is a Canadian computer scientist, best known as the founder and lead designer behind the Java programming language. https://en.wikipedia.org/wiki/James_Gosling
		"gosling",

		// Stephen Jay Gould was was an American paleontologist, evolutionary biologist, and historian of science. He is most famous for the theory of punctuated equilibrium - https://en.wikipedia.org/wiki/Stephen_Jay_Gould
		"gould",

		// Carolyn Widney Greider - American molecular biologist and joint winner of the 2009 Nobel Prize for Physiology or Medicine for the discovery of telomerase. https://en.wikipedia.org/wiki/Carol_W._Greider
		"greider",

		// Alexander Grothendieck - German-born French mathematician who became a leading figure in the creation of modern algebraic geometry. https://en.wikipedia.org/wiki/Alexander_Grothendieck
		"grothendieck",

		// Lois Haibt - American computer scientist, part of the team at IBM that developed FORTRAN - https://en.wikipedia.org/wiki/Lois_Haibt
		"haibt",

		// Margaret Hamilton - Director of the Software Engineering Division of the MIT Instrumentation Laboratory, which developed on-board flight software for the Apollo space program. https://en.wikipedia.org/wiki/Margaret_Hamilton_(scientist)
		"hamilton",

		// Caroline Harriet Haslett - English electrical engineer, electricity industry administrator and champion of women's rights. Co-author of British Standard 1363 that specifies AC power plugs and sockets used across the United Kingdom (which is widely considered as one of the safest designs). https://en.wikipedia.org/wiki/Caroline_Haslett
		"haslett",

		// Stephen Hawking pioneered the field of cosmology by combining general relativity and quantum mechanics. https://en.wikipedia.org/wiki/Stephen_Hawking
		"hawking",

		// Martin Edward Hellman - American cryptologist, best known for his invention of public-key cryptography in co-operation with Whitfield Diffie and Ralph Merkle. https://en.wikipedia.org/wiki/Martin_Hellman
		"hellman",

		// Werner Heisenberg was a founding father of quantum mechanics. https://en.wikipedia.org/wiki/Werner_Heisenberg
		"heisenberg",

		// Grete Hermann was a German philosopher noted for her philosophical work on the foundations of quantum mechanics. https://en.wikipedia.org/wiki/Grete_Hermann
		"hermann",

		// Caroline Lucretia Herschel - German astronomer and discoverer of several comets. https://en.wikipedia.org/wiki/Caroline_Herschel
		"herschel",

		// Heinrich Rudolf Hertz - German physicist who first conclusively proved the existence of the electromagnetic waves. https://en.wikipedia.org/wiki/Heinrich_Hertz
		"hertz",

		// Jaroslav Heyrovský was the inventor of the polarographic method, father of the electroanalytical method, and recipient of the Nobel Prize in 1959. His main field of work was polarography. https://en.wikipedia.org/wiki/Jaroslav_Heyrovsk%C3%BD
		"heyrovsky",

		// Dorothy Hodgkin was a British biochemist, credited with the development of protein crystallography. She was awarded the Nobel Prize in Chemistry in 1964. https://en.wikipedia.org/wiki/Dorothy_Hodgkin
		"hodgkin",

		// Douglas R. Hofstadter is an American professor of cognitive science and author of the Pulitzer Prize and American Book Award-winning work Goedel, Escher, Bach: An Eternal Golden Braid in 1979. A mind-bending work which coined Hofstadter's Law: "It always takes longer than you expect, even when you take into account Hofstadter's Law." https://en.wikipedia.org/wiki/Douglas_Hofstadter
		"hofstadter",

		// Erna Schneider Hoover revolutionized modern communication by inventing a computerized telephone switching method. https://en.wikipedia.org/wiki/Erna_Schneider_Hoover
		"hoover",

		// Grace Hopper developed the first compiler for a computer programming language and  is credited with popularizing the term "debugging" for fixing computer glitches. https://en.wikipedia.org/wiki/Grace_Hopper
		"hopper",

		// Frances Hugle, she was an American scientist, engineer, and inventor who contributed to the understanding of semiconductors, integrated circuitry, and the unique electrical principles of microscopic materials. https://en.wikipedia.org/wiki/Frances_Hugle
		"hugle",

		// Hypatia - Greek Alexandrine Neoplatonist philosopher in Egypt who was one of the earliest mothers of mathematics - https://en.wikipedia.org/wiki/Hypatia
		"hypatia",

		// Teruko Ishizaka - Japanese scientist and immunologist who co-discovered the antibody class Immunoglobulin E. https://en.wikipedia.org/wiki/Teruko_Ishizaka
		"ishizaka",

		// Mary Jackson, American mathematician and aerospace engineer who earned the highest title within NASA's engineering department - https://en.wikipedia.org/wiki/Mary_Jackson_(engineer)
		"jackson",

		// Yeong-Sil Jang was a Korean scientist and astronomer during the Joseon Dynasty; he invented the first metal printing press and water gauge. https://en.wikipedia.org/wiki/Jang_Yeong-sil
		"jang",

		// Frederick Jelinek - Czech-American scientist, father of statistical automatic speech recognition. https://en.wikipedia.org/wiki/Frederick_Jelinek
		"jelinek",

		// Mae Carol Jemison -  is an American engineer, physician, and former NASA astronaut. She became the first black woman to travel in space when she served as a mission specialist aboard the Space Shuttle Endeavour - https://en.wikipedia.org/wiki/Mae_Jemison
		"jemison",

		// Betty Jennings - one of the original programmers of the ENIAC. https://en.wikipedia.org/wiki/ENIAC - https://en.wikipedia.org/wiki/Jean_Bartik
		"jennings",

		// Mary Lou Jepsen, was the founder and chief technology officer of One Laptop Per Child (OLPC), and the founder of Pixel Qi. https://en.wikipedia.org/wiki/Mary_Lou_Jepsen
		"jepsen",

		// Katherine Coleman Goble Johnson - American physicist and mathematician contributed to the NASA. https://en.wikipedia.org/wiki/Katherine_Johnson
		"johnson",

		// Irène Joliot-Curie - French scientist who was awarded the Nobel Prize for Chemistry in 1935. Daughter of Marie and Pierre Curie. https://en.wikipedia.org/wiki/Ir%C3%A8ne_Joliot-Curie
		"joliot",

		// Karen Spärck Jones came up with the concept of inverse document frequency, which is used in most search engines today. https://en.wikipedia.org/wiki/Karen_Sp%C3%A4rck_Jones
		"jones",

		// A. P. J. Abdul Kalam - is an Indian scientist aka Missile Man of India for his work on the development of ballistic missile and launch vehicle technology - https://en.wikipedia.org/wiki/A._P._J._Abdul_Kalam
		"kalam",

		// Sergey Petrovich Kapitsa (Russian: Серге́й Петро́вич Капи́ца; 14 February 1928 – 14 August 2012) was a Russian physicist and demographer. He was best known as host of the popular and long-running Russian scientific TV show, Evident, but Incredible. His father was the Nobel laureate Soviet-era physicist Pyotr Kapitsa, and his brother was the geographer and Antarctic explorer Andrey Kapitsa. - https://en.wikipedia.org/wiki/Sergey_Kapitsa
		"kapitsa",

		// Susan Kare, created the icons and many of the interface elements for the original Apple Macintosh in the 1980s, and was an original employee of NeXT, working as the Creative Director. https://en.wikipedia.org/wiki/Susan_Kare
		"kare",

		// Mstislav Keldysh - a Soviet scientist in the field of mathematics and mechanics, academician of the USSR Academy of Sciences (1946), President of the USSR Academy of Sciences (1961–1975), three times Hero of Socialist Labor (1956, 1961, 1971), fellow of the Royal Society of Edinburgh (1968). https://en.wikipedia.org/wiki/Mstislav_Keldysh
		"keldysh",

		// Mary Kenneth Keller, Sister Mary Kenneth Keller became the first American woman to earn a PhD in Computer Science in 1965. https://en.wikipedia.org/wiki/Mary_Kenneth_Keller
		"keller",

		// Johannes Kepler, German astronomer known for his three laws of planetary motion - https://en.wikipedia.org/wiki/Johannes_Kepler
		"kepler",

		// Omar Khayyam - Persian mathematician, astronomer and poet. Known for his work on the classification and solution of cubic equations, for his contribution to the understanding of Euclid's fifth postulate and for computing the length of a year very accurately. https://en.wikipedia.org/wiki/Omar_Khayyam
		"khayyam",

		// Har Gobind Khorana - Indian-American biochemist who shared the 1968 Nobel Prize for Physiology - https://en.wikipedia.org/wiki/Har_Gobind_Khorana
		"khorana",

		// Jack Kilby invented silicon integrated circuits and gave Silicon Valley its name. - https://en.wikipedia.org/wiki/Jack_Kilby
		"kilby",

		// Maria Kirch - German astronomer and first woman to discover a comet - https://en.wikipedia.org/wiki/Maria_Margarethe_Kirch
		"kirch",

		// Donald Knuth - American computer scientist, author of "The Art of Computer Programming" and creator of the TeX typesetting system. https://en.wikipedia.org/wiki/Donald_Knuth
		"knuth",

		// Sophie Kowalevski - Russian mathematician responsible for important original contributions to analysis, differential equations and mechanics - https://en.wikipedia.org/wiki/Sofia_Kovalevskaya
		"kowalevski",

		// Stephanie Louise Kwolek - American chemist who invented Kevlar. https://en.wikipedia.org/wiki/Stephanie_Kwolek
		"kwolek",

		// Marie-Jeanne de Lalande - French astronomer, mathematician and cataloguer of stars - https://en.wikipedia.org/wiki/Marie-Jeanne_de_Lalande
		"lalande",

		// Hedy Lamarr - Actress and inventor. The principles of her work are now incorporated into modern Wi-Fi, CDMA and Bluetooth technology. https://en.wikipedia.org/wiki/Hedy_Lamarr
		"lamarr",

		// Leslie B. Lamport - American computer scientist. Lamport is best known for his seminal work in distributed systems and was the winner of the 2013 Turing Award. https://en.wikipedia.org/wiki/Leslie_Lamport
		"lamport",

		// Mary Leakey - British paleoanthropologist who discovered the first fossilized Proconsul skull - https://en.wikipedia.org/wiki/Mary_Leakey
		"leakey",

		// Henrietta Swan Leavitt - she was an American astronomer who discovered the relation between the luminosity and the period of Cepheid variable stars. https://en.wikipedia.org/wiki/Henrietta_Swan_Leavitt
		"leavitt",

		// Esther Miriam Zimmer Lederberg - American microbiologist and a pioneer of bacterial genetics. https://en.wikipedia.org/wiki/Esther_Lederberg
		"lederberg",

		// Inge Lehmann - Danish seismologist and geophysicist. Known for discovering in 1936 that the Earth has a solid inner core inside a molten outer core. https://en.wikipedia.org/wiki/Inge_Lehmann
		"lehmann",

		// Leucippus - Believed to be the founder of the Atomist school of philosophy. Mentor to Democritus. https://en.wikipedia.org/wiki/Leucippus
		"leucippus",

		// Vladimir Levenshtein - Russian scientist known for his work on combinatorial design, error-correcting codes and information theory; creator of the Levenshtein distance. https://en.wikipedia.org/wiki/Vladimir_Levenshtein
		"levenshtein",

		// Daniel Lewin - Mathematician, Akamai co-founder, soldier, 9/11 victim-- Developed optimization techniques for routing traffic on the internet. Died attempting to stop the 9-11 hijackers. https://en.wikipedia.org/wiki/Daniel_Lewin
		"lewin",

		// Ruth Lichterman - one of the original programmers of the ENIAC. https://en.wikipedia.org/wiki/ENIAC - https://en.wikipedia.org/wiki/Ruth_Teitelbaum
		"lichterman",

		// Barbara Liskov - co-developed the Liskov substitution principle. Liskov was also the winner of the Turing Prize in 2008. - https://en.wikipedia.org/wiki/Barbara_Liskov
		"liskov",

		// Ada Lovelace invented the first algorithm. https://en.wikipedia.org/wiki/Ada_Lovelace (thanks James Turnbull)
		"lovelace",

		// Lucretius - ancient greek philosopher whose ideas about the natural world were early precursors to Darwin's Theory of Evolution. Also a proponent of Atomism and Empiricism. https://en.wikipedia.org/wiki/Lucretius
		"lucretius",

		// Auguste and Louis Lumière - the first filmmakers in history - https://en.wikipedia.org/wiki/Auguste_and_Louis_Lumi%C3%A8re
		"lumiere",

		// Prasanta Chandra Mahalanobis - Indian applied statistician, best known for the Mahalanobis distance and for making ground-breaking contributions related to large-scale sample surveys. https://en.wikipedia.org/wiki/Prasanta_Chandra_Mahalanobis
		"mahalanobis",

		// Mahavira - Ancient Indian mathematician during 9th century AD who discovered basic algebraic identities - https://en.wikipedia.org/wiki/Mah%C4%81v%C4%ABra_(mathematician)
		"mahavira",

		// Lynn Margulis (b. Lynn Petra Alexander) - an American evolutionary theorist and biologist, science author, educator, and popularizer, and was the primary modern proponent for the significance of symbiosis in evolution. - https://en.wikipedia.org/wiki/Lynn_Margulis
		"margulis",

		// Yukihiro Matsumoto - Japanese computer scientist and software programmer best known as the chief designer of the Ruby programming language. https://en.wikipedia.org/wiki/Yukihiro_Matsumoto
		"matsumoto",

		// James Clerk Maxwell - Scottish physicist, best known for his formulation of electromagnetic theory. https://en.wikipedia.org/wiki/James_Clerk_Maxwell
		"maxwell",

		// Maria Mayer - American theoretical physicist and Nobel laureate in Physics for proposing the nuclear shell model of the atomic nucleus - https://en.wikipedia.org/wiki/Maria_Mayer
		"mayer",

		// John McCarthy invented LISP: https://en.wikipedia.org/wiki/John_McCarthy_(computer_scientist)
		"mccarthy",

		// Barbara McClintock - a distinguished American cytogeneticist, 1983 Nobel Laureate in Physiology or Medicine for discovering transposons. https://en.wikipedia.org/wiki/Barbara_McClintock
		"mcclintock",

		// Anne Laura Dorinthea McLaren - British developmental biologist whose work helped lead to human in-vitro fertilisation. https://en.wikipedia.org/wiki/Anne_McLaren
		"mclaren",

		// Malcolm McLean invented the modern shipping container: https://en.wikipedia.org/wiki/Malcom_McLean
		"mclean",

		// Kay McNulty - one of the original programmers of the ENIAC. https://en.wikipedia.org/wiki/ENIAC - https://en.wikipedia.org/wiki/Kathleen_Antonelli
		"mcnulty",

		// Gregor Johann Mendel - Czech scientist and founder of genetics. https://en.wikipedia.org/wiki/Gregor_Mendel
		"mendel",

		// Dmitri Mendeleev - a chemist and inventor. He formulated the Periodic Law, created a farsighted version of the periodic table of elements, and used it to correct the properties of some already discovered elements and also to predict the properties of eight elements yet to be discovered. https://en.wikipedia.org/wiki/Dmitri_Mendeleev
		"mendeleev",

		// Lise Meitner - Austrian/Swedish physicist who was involved in the discovery of nuclear fission. The element meitnerium is named after her - https://en.wikipedia.org/wiki/Lise_Meitner
		"meitner",

		// Carla Meninsky, was the game designer and programmer for Atari 2600 games Dodge 'Em and Warlords. https://en.wikipedia.org/wiki/Carla_Meninsky
		"meninsky",

		// Ralph C. Merkle - American computer scientist, known for devising Merkle's puzzles - one of the very first schemes for public-key cryptography. Also, inventor of Merkle trees and co-inventor of the Merkle-Damgård construction for building collision-resistant cryptographic hash functions and the Merkle-Hellman knapsack cryptosystem. https://en.wikipedia.org/wiki/Ralph_Merkle
		"merkle",

		// Johanna Mestorf - German prehistoric archaeologist and first female museum director in Germany - https://en.wikipedia.org/wiki/Johanna_Mestorf
		"mestorf",

		// Maryam Mirzakhani - an Iranian mathematician and the first woman to win the Fields Medal. https://en.wikipedia.org/wiki/Maryam_Mirzakhani
		"mirzakhani",

		// Maria Mitchell - first female professional astronomer in the US, discoverer of "Miss Mitchell's Comet". https://en.wikipedia.org/wiki/Maria_Mitchell
		"mitchell",

		// Rita Levi-Montalcini - Won Nobel Prize in Physiology or Medicine jointly with colleague Stanley Cohen for the discovery of nerve growth factor (https://en.wikipedia.org/wiki/Rita_Levi-Montalcini)
		"montalcini",

		// Rita Levi-Montalcini - Won Nobel Prize in Physiology or Medicine jointly with colleague Stanley Cohen for the discovery of nerve growth factor (https://en.wikipedia.org/wiki/Rita_Levi-Montalcini)
		"montalcini",

		// Gordon Earle Moore - American engineer, Silicon Valley founding father, author of Moore's law. https://en.wikipedia.org/wiki/Gordon_Moore
		"moore",

		// Mary Sherman Morgan - American rocket fuel scientist who invented the liquid rocket fuel Hydyne. https://en.wikipedia.org/wiki/Mary_Sherman_Morgan
		"morgan",

		// Samuel Morse - contributed to the invention of a single-wire telegraph system based on European telegraphs and was a co-developer of the Morse code - https://en.wikipedia.org/wiki/Samuel_Morse
		"morse",

		// Ian Murdock - founder of the Debian project - https://en.wikipedia.org/wiki/Ian_Murdock
		"murdock",

		// May-Britt Moser - Nobel prize winner neuroscientist who contributed to the discovery of grid cells in the brain. https://en.wikipedia.org/wiki/May-Britt_Moser
		"moser",

		// John Napier of Merchiston - Scottish landowner known as an astronomer, mathematician and physicist. Best known for his discovery of logarithms. https://en.wikipedia.org/wiki/John_Napier
		"napier",

		// John Forbes Nash, Jr. - American mathematician who made fundamental contributions to game theory, differential geometry, and the study of partial differential equations. https://en.wikipedia.org/wiki/John_Forbes_Nash_Jr.
		"nash",

		// John von Neumann - todays computer architectures are based on the von Neumann architecture. https://en.wikipedia.org/wiki/Von_Neumann_architecture
		"neumann",

		// Isaac Newton invented classic mechanics and modern optics. https://en.wikipedia.org/wiki/Isaac_Newton
		"newton",

		// Florence Nightingale, more prominently known as a nurse, was also the first female member of the Royal Statistical Society and a pioneer in statistical graphics https://en.wikipedia.org/wiki/Florence_Nightingale#Statistics_and_sanitary_reform
		"nightingale",

		// Alfred Nobel - a Swedish chemist, engineer, innovator, and armaments manufacturer (inventor of dynamite) - https://en.wikipedia.org/wiki/Alfred_Nobel
		"nobel",

		// Ida Noddack - German chemist and physicist, best known for suggesting the idea of nuclear fission for the first time and for the discovery of rhenium. https://en.wikipedia.org/wiki/Ida_Noddack
		"noddack",

		// Emmy Noether, German mathematician. Noether's Theorem is named after her. https://en.wikipedia.org/wiki/Emmy_Noether
		"noether",

		// Poppy Northcutt. Poppy Northcutt was the first woman to work as part of NASA’s Mission Control. http://www.businessinsider.com/poppy-northcutt-helped-apollo-astronauts-2014-12?op=1
		"northcutt",

		// Robert Noyce invented silicon integrated circuits and gave Silicon Valley its name. - https://en.wikipedia.org/wiki/Robert_Noyce
		"noyce",

		// Julius Robert Oppenheimer - American theoretical physicist best known for the Born–Oppenheimer approximation for molecular wave functions, his work on the theory of electrons and positrons, the Oppenheimer–Phillips process in nuclear fusion, the first prediction of quantum tunneling and as the leader of the Manhattan Project. https://en.wikipedia.org/wiki/J._Robert_Oppenheimer
		"oppenheimer",

		// Panini - Ancient Indian linguist and grammarian from 4th century CE who worked on the world's first formal system - https://en.wikipedia.org/wiki/P%C4%81%E1%B9%87ini#Comparison_with_modern_formal_systems
		"panini",

		// George Papanicolaou - Greek pioneer in cytopathology and early cancer detection; inventor of the Pap smear. https://en.wikipedia.org/wiki/Georgios_Papanikolaou
		"papanicolaou",

		// Ambroise Pare invented modern surgery. https://en.wikipedia.org/wiki/Ambroise_Par%C3%A9
		"pare",

		// Blaise Pascal, French mathematician, physicist, and inventor - https://en.wikipedia.org/wiki/Blaise_Pascal
		"pascal",

		// Louis Pasteur discovered vaccination, fermentation and pasteurization. https://en.wikipedia.org/wiki/Louis_Pasteur.
		"pasteur",

		// Cecilia Payne-Gaposchkin was an astronomer and astrophysicist who, in 1925, proposed in her Ph.D. thesis an explanation for the composition of stars in terms of the relative abundances of hydrogen and helium. https://en.wikipedia.org/wiki/Cecilia_Payne-Gaposchkin
		"payne",

		// Radia Perlman is a software designer and network engineer and most famous for her invention of the spanning-tree protocol (STP). https://en.wikipedia.org/wiki/Radia_Perlman
		"perlman",

		// Rob Pike was a key contributor to Unix, Plan 9, the X graphic system, utf-8, and the Go programming language. https://en.wikipedia.org/wiki/Rob_Pike
		"pike",

		// Henri Poincaré made fundamental contributions in several fields of mathematics. https://en.wikipedia.org/wiki/Henri_Poincar%C3%A9
		"poincare",

		// Laura Poitras is a director and producer whose work, made possible by open source crypto tools, advances the causes of truth and freedom of information by reporting disclosures by whistleblowers such as Edward Snowden. https://en.wikipedia.org/wiki/Laura_Poitras
		"poitras",

		// Tat’yana Avenirovna Proskuriakova (Russian: Татья́на Авени́ровна Проскуряко́ва) (January 23 [O.S. January 10] 1909 – August 30, 1985) was a Russian-American Mayanist scholar and archaeologist who contributed significantly to the deciphering of Maya hieroglyphs, the writing system of the pre-Columbian Maya civilization of Mesoamerica. https://en.wikipedia.org/wiki/Tatiana_Proskouriakoff
		"proskuriakova",

		// Claudius Ptolemy - a Greco-Egyptian writer of Alexandria, known as a mathematician, astronomer, geographer, astrologer, and poet of a single epigram in the Greek Anthology - https://en.wikipedia.org/wiki/Ptolemy
		"ptolemy",

		// C. V. Raman - Indian physicist who won the Nobel Prize in 1930 for proposing the Raman effect. - https://en.wikipedia.org/wiki/C._V._Raman
		"raman",

		// Srinivasa Ramanujan - Indian mathematician and autodidact who made extraordinary contributions to mathematical analysis, number theory, infinite series, and continued fractions. - https://en.wikipedia.org/wiki/Srinivasa_Ramanujan
		"ramanujan",

		// Marian Adam Rejewski - Polish mathematician and cryptologist who reconstructed the Nazi German Enigma cipher machine sight-unseen in 1932. https://en.wikipedia.org/wiki/Marian_Rejewski
		"rejewski",

		// Sally Kristen Ride was an American physicist and astronaut. She was the first American woman in space, and the youngest American astronaut. https://en.wikipedia.org/wiki/Sally_Ride
		"ride",

		// Dennis Ritchie - co-creator of UNIX and the C programming language. - https://en.wikipedia.org/wiki/Dennis_Ritchie
		"ritchie",

		// Ida Rhodes - American pioneer in computer programming, designed the first computer used for Social Security. https://en.wikipedia.org/wiki/Ida_Rhodes
		"rhodes",

		// Julia Hall Bowman Robinson - American mathematician renowned for her contributions to the fields of computability theory and computational complexity theory. https://en.wikipedia.org/wiki/Julia_Robinson
		"robinson",

		// Wilhelm Conrad Röntgen - German physicist who was awarded the first Nobel Prize in Physics in 1901 for the discovery of X-rays (Röntgen rays). https://en.wikipedia.org/wiki/Wilhelm_R%C3%B6ntgen
		"roentgen",

		// Rosalind Franklin - British biophysicist and X-ray crystallographer whose research was critical to the understanding of DNA - https://en.wikipedia.org/wiki/Rosalind_Franklin
		"rosalind",

		// Jerzy Witold Rozycki - Polish mathematician and cryptologist who (together with Marian Rejewski and Henryk Zygalski) solved the German military Enigma cipher machine. https://en.wikipedia.org/wiki/Jerzy_R%C3%B3%C5%BCycki
		"rozycki",

		// Vera Rubin - American astronomer who pioneered work on galaxy rotation rates. https://en.wikipedia.org/wiki/Vera_Rubin
		"rubin",

		// Joanna Rutkowska - Polish computer security researcher. https://en.wikipedia.org/wiki/Joanna_Rutkowska
		"rutkowska",

		// Meghnad Saha - Indian astrophysicist best known for his development of the Saha equation, used to describe chemical and physical conditions in stars - https://en.wikipedia.org/wiki/Meghnad_Saha
		"saha",

		// Jean E. Sammet developed FORMAC, the first widely used computer language for symbolic manipulation of mathematical formulas. https://en.wikipedia.org/wiki/Jean_E._Sammet
		"sammet",

		// Mildred Sanderson - American mathematician best known for Sanderson's theorem concerning modular invariants. https://en.wikipedia.org/wiki/Mildred_Sanderson
		"sanderson",

		// Katsuko Saruhashi - Japanese geochemist known for her creation of the first methods for measuring carbon dioxide levels in seawater. She also developed a technique to trace the travel of radioactive fallout across the oceans that led to restricting oceanic nuclear experimentation in 1963. https://en.wikipedia.org/wiki/Katsuko_Saruhashi
		"saruhashi",

		// Satoshi Nakamoto is the name used by the unknown person or group of people who developed bitcoin, authored the bitcoin white paper, and created and deployed bitcoin's original reference implementation. https://en.wikipedia.org/wiki/Satoshi_Nakamoto
		"satoshi",

		// Adi Shamir - Israeli cryptographer whose numerous inventions and contributions to cryptography include the Ferge Fiat Shamir identification scheme, the Rivest Shamir Adleman (RSA) public-key cryptosystem, the Shamir's secret sharing scheme, the breaking of the Merkle-Hellman cryptosystem, the TWINKLE and TWIRL factoring devices and the discovery of differential cryptanalysis (with Eli Biham). https://en.wikipedia.org/wiki/Adi_Shamir
		"shamir",

		// Claude Shannon - The father of information theory and founder of digital circuit design theory. (https://en.wikipedia.org/wiki/Claude_Shannon)
		"shannon",

		// Carol Shaw - Originally an Atari employee, Carol Shaw is said to be the first female video game designer. https://en.wikipedia.org/wiki/Carol_Shaw_(video_game_designer)
		"shaw",

		// Dame Stephanie "Steve" Shirley - Founded a software company in 1962 employing women working from home. https://en.wikipedia.org/wiki/Steve_Shirley
		"shirley",

		// William Shockley co-invented the transistor - https://en.wikipedia.org/wiki/William_Shockley
		"shockley",

		// Lina Solomonovna Stern (or Shtern; Russian: Лина Соломоновна Штерн; 26 August 1878 – 7 March 1968) was a Soviet biochemist, physiologist and humanist whose medical discoveries saved thousands of lives at the fronts of World War II. She is best known for her pioneering work on blood–brain barrier, which she described as hemato-encephalic barrier in 1921. https://en.wikipedia.org/wiki/Lina_Stern
		"shtern",

		// Françoise Barré-Sinoussi - French virologist and Nobel Prize Laureate in Physiology or Medicine; her work was fundamental in identifying HIV as the cause of AIDS. https://en.wikipedia.org/wiki/Fran%C3%A7oise_Barr%C3%A9-Sinoussi
		"sinoussi",

		// Betty Snyder - one of the original programmers of the ENIAC. https://en.wikipedia.org/wiki/ENIAC - https://en.wikipedia.org/wiki/Betty_Holberton
		"snyder",

		// Cynthia Solomon - Pioneer in the fields of artificial intelligence, computer science and educational computing. Known for creation of Logo, an educational programming language.  https://en.wikipedia.org/wiki/Cynthia_Solomon
		"solomon",

		// Frances Spence - one of the original programmers of the ENIAC. https://en.wikipedia.org/wiki/ENIAC - https://en.wikipedia.org/wiki/Frances_Spence
		"spence",

		// Richard Matthew Stallman - the founder of the Free Software movement, the GNU project, the Free Software Foundation, and the League for Programming Freedom. He also invented the concept of copyleft to protect the ideals of this movement, and enshrined this concept in the widely-used GPL (General Public License) for software. https://en.wikiquote.org/wiki/Richard_Stallman
		"stallman",

		// Nettie Maria Stevens - American geneticist credited with the discovery of sex chromosomes. https://en.wikipedia.org/wiki/Nettie_Stevens
		"stevens",

		// Michael Stonebraker is a database research pioneer and architect of Ingres, Postgres, VoltDB and SciDB. Winner of 2014 ACM Turing Award. https://en.wikipedia.org/wiki/Michael_Stonebraker
		"stonebraker",

		// Bjarne Stroustrup - Creator of C++. https://en.wikipedia.org/wiki/Bjarne_Stroustrup
		"stroustrup",

		// Ivan Edward Sutherland - American computer scientist and Internet pioneer, widely regarded as the father of computer graphics. https://en.wikipedia.org/wiki/Ivan_Sutherland
		"sutherland",

		// Janese Swanson (with others) developed the first of the Carmen Sandiego games. She went on to found Girl Tech. https://en.wikipedia.org/wiki/Janese_Swanson
		"swanson",

		// Aaron Swartz was influential in creating RSS, Markdown, Creative Commons, Reddit, and much of the internet as we know it today. He was devoted to freedom of information on the web. https://en.wikiquote.org/wiki/Aaron_Swartz
		"swartz",

		// Bertha Swirles was a theoretical physicist who made a number of contributions to early quantum theory. https://en.wikipedia.org/wiki/Bertha_Swirles
		"swirles",

		// Robert Endre Tarjan - American computer scientist and mathematician. Inventor of several graph algorithms, including Tarjan's off-line lowest common ancestors algorithm, and co-inventor of Fibonacci heaps and splay trees. https://en.wikipedia.org/wiki/Robert_Tarjan
		"tarjan",

		// Helen Brooke Taussig - American cardiologist and founder of the field of paediatric cardiology. https://en.wikipedia.org/wiki/Helen_B._Taussig
		"taussig",

		// Valentina Tereshkova is a Russian engineer, cosmonaut and politician. She was the first woman to fly to space in 1963. In 2013, at the age of 76, she offered to go on a one-way mission to Mars. https://en.wikipedia.org/wiki/Valentina_Tereshkova
		"tereshkova",

		// Nikola Tesla invented the AC electric system and every gadget ever used by a James Bond villain. https://en.wikipedia.org/wiki/Nikola_Tesla
		"tesla",

		// Marie Tharp - American geologist and oceanic cartographer who co-created the first scientific map of the Atlantic Ocean floor. Her work led to the acceptance of the theories of plate tectonics and continental drift. https://en.wikipedia.org/wiki/Marie_Tharp
		"tharp",

		// Ken Thompson - co-creator of UNIX and the C programming language - https://en.wikipedia.org/wiki/Ken_Thompson
		"thompson",

		// Linus Torvalds invented Linux and Git. https://en.wikipedia.org/wiki/Linus_Torvalds
		"torvalds",

		// Youyou Tu - Chinese pharmaceutical chemist and educator known for discovering artemisinin and dihydroartemisinin, used to treat malaria, which has saved millions of lives. Joint winner of the 2015 Nobel Prize in Physiology or Medicine. https://en.wikipedia.org/wiki/Tu_Youyou
		"tu",

		// Alan Turing was a founding father of computer science. https://en.wikipedia.org/wiki/Alan_Turing.
		"turing",

		// Varahamihira - Ancient Indian mathematician who discovered trigonometric formulae during 505-587 CE - https://en.wikipedia.org/wiki/Var%C4%81hamihira#Contributions
		"varahamihira",

		// Dorothy Vaughan was a NASA mathematician and computer programmer on the SCOUT launch vehicle program that put America's first satellites into space - https://en.wikipedia.org/wiki/Dorothy_Vaughan
		"vaughan",

		// Cédric Villani - French mathematician, won Fields Medal, Fermat Prize and Poincaré Price for his work in differential geometry and statistical mechanics. https://en.wikipedia.org/wiki/C%C3%A9dric_Villani
		"villani",

		// Artturi Ilmari Virtanen - Finnish chemist and recipient of the 1945 Nobel Prize in Chemistry "for his research and inventions in agricultural and nutrition chemistry, especially for his fodder preservation method". https://en.wikipedia.org/wiki/Artturi_Ilmari_Virtanen
		"virtanen",

		// Sir Mokshagundam Visvesvaraya - is a notable Indian engineer.  He is a recipient of the Indian Republic's highest honour, the Bharat Ratna, in 1955. On his birthday, 15 September is celebrated as Engineer's Day in India in his memory - https://en.wikipedia.org/wiki/Visvesvaraya
		"visvesvaraya",

		// Christiane Nüsslein-Volhard - German biologist, won Nobel Prize in Physiology or Medicine in 1995 for research on the genetic control of embryonic development. https://en.wikipedia.org/wiki/Christiane_N%C3%BCsslein-Volhard
		"volhard",

		// Marlyn Wescoff - one of the original programmers of the ENIAC. https://en.wikipedia.org/wiki/ENIAC - https://en.wikipedia.org/wiki/Marlyn_Meltzer
		"wescoff",

		// Sylvia B. Wilbur - British computer scientist who helped develop the ARPANET, was one of the first to exchange email in the UK and a leading researcher in computer-supported collaborative work. https://en.wikipedia.org/wiki/Sylvia_Wilbur
		"wilbur",

		// Andrew Wiles - Notable British mathematician who proved the enigmatic Fermat's Last Theorem - https://en.wikipedia.org/wiki/Andrew_Wiles
		"wiles",

		// Roberta Williams, did pioneering work in graphical adventure games for personal computers, particularly the King's Quest series. https://en.wikipedia.org/wiki/Roberta_Williams
		"williams",

		// Malcolm John Williamson - British mathematician and cryptographer employed by the GCHQ. Developed in 1974 what is now known as Diffie-Hellman key exchange (Diffie and Hellman first published the scheme in 1976). https://en.wikipedia.org/wiki/Malcolm_J._Williamson
		"williamson",

		// Lucy Wills - English haematologist whose analysis of prenatal anemia changed the face of preventive prenatal care for women everywhere. https://en.wikipedia.org/wiki/Lucy_Wills
		"wills",

		// Sophie Wilson designed the first Acorn Micro-Computer and the instruction set for ARM processors. https://en.wikipedia.org/wiki/Sophie_Wilson
		"wilson",

		// Jeannette Wing - co-developed the Liskov substitution principle. - https://en.wikipedia.org/wiki/Jeannette_Wing
		"wing",

		// Steve Wozniak invented the Apple I and Apple II. https://en.wikipedia.org/wiki/Steve_Wozniak
		"wozniak",

		// The Wright brothers, Orville and Wilbur - credited with inventing and building the world's first successful airplane and making the first controlled, powered and sustained heavier-than-air human flight - https://en.wikipedia.org/wiki/Wright_brothers
		"wright",

		// Chien-Shiung Wu - Chinese-American experimental physicist who made significant contributions to nuclear physics. https://en.wikipedia.org/wiki/Chien-Shiung_Wu
		"wu",

		// Rosalyn Sussman Yalow - Rosalyn Sussman Yalow was an American medical physicist, and a co-winner of the 1977 Nobel Prize in Physiology or Medicine for development of the radioimmunoassay technique. https://en.wikipedia.org/wiki/Rosalyn_Sussman_Yalow
		"yalow",

		// Ada Yonath - an Israeli crystallographer, the first woman from the Middle East to win a Nobel prize in the sciences. https://en.wikipedia.org/wiki/Ada_Yonath
		"yonath",

		// Lotfi Aliasker Zadeh - Azerbaijani scientist best known for proposing fuzzy mathematics. https://en.wikipedia.org/wiki/Lotfi_A._Zadeh
		"zadeh",

		// Yaakov Ziv - Israeli electrical engineer who, along with Abraham Lempel, developed the LZ family of lossless data compression algorithms. https://en.wikipedia.org/wiki/Yaakov_Ziv
		"ziv",

		// Nikolay Yegorovich Zhukovsky (Russian: Никола́й Его́рович Жуко́вский, January 17 1847 – March 17, 1921) was a Russian scientist, mathematician and engineer, and a founding father of modern aero- and hydrodynamics. Whereas contemporary scientists scoffed at the idea of human flight, Zhukovsky was the first to undertake the study of airflow. He is often called the Father of Russian Aviation. https://en.wikipedia.org/wiki/Nikolay_Yegorovich_Zhukovsky
		"zhukovsky",

		// Henryk Zygalski - Polish mathematician and cryptologist who (together with Marian Rejewski and Jerzy Rozycki) solved the German military Enigma cipher machine. https://en.wikipedia.org/wiki/Henryk_Zygalski
		"zygalski",
	}
)

var localRand = rand.New(rand.NewSource(time.Now().UnixNano()))
var leftLen = len(left)
var rightLen = len(right)

// GetNameWithSeed generates a random name from the list of adjectives and surnames in this package
// formatted as "adjective_surname". For example 'focused_turing'. If retry is non-zero, a random
// integer between 0 and 10 will be added to the end of the name, e.g `focused_turing3`
func GetNameWithSeed(r *rand.Rand, retry int) string {
	if r == nil {
		r = localRand
	}
begin:
	name := left[r.Intn(leftLen)] + "_" + right[r.Intn(rightLen)] //nolint:gosec // G404: Use of weak random number generator (math/rand instead of crypto/rand)
	if name == "boring_wozniak" /* Steve Wozniak is not boring */ {
		goto begin
	}
	if retry > 0 {
		name += strconv.Itoa(r.Intn(10)) //nolint:gosec // G404: Use of weak random number generator (math/rand instead of crypto/rand)
	}
	return name
}

// GetName generates a random name from the list of adjectives and surnames in this package
// formatted as "adjective_surname". For example 'focused_turing'. If retry is non-zero, a random
// integer between 0 and 10 will be added to the end of the name, e.g `focused_turing3`
func GetName(retry int) string {
begin:
	name := left[rand.Intn(leftLen)] + "_" + right[rand.Intn(rightLen)] //nolint:gosec // G404: Use of weak random number generator (math/rand instead of crypto/rand)
	if name == "boring_wozniak" /* Steve Wozniak is not boring */ {
		goto begin
	}
	if retry > 0 {
		name += strconv.Itoa(rand.Intn(10)) //nolint:gosec // G404: Use of weak random number generator (math/rand instead of crypto/rand)
	}
	return name
}
