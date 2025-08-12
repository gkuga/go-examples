class Once
  def initialize
    @done = false
    @m = Mutex.new
  end

  def do
    return if @done
    do_slow { yield }
  end

  private

  def do_slow
    @m.synchronize do
      return if @done
      begin
        yield
      ensure
        @done = true
      end
    end
  end
end


once = Once.new
threads = 10.times.map { |i|
  Thread.new { 
    once.do {
      puts "run once"
    }
    print i
  }
}
threads.each(&:join)
